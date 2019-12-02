package driver

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/ucloud/ucloud-sdk-go/internal/utest"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	uerr "github.com/ucloud/ucloud-sdk-go/ucloud/error"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
	"testing"
	"time"
)

type StepReport struct {
	Title      string        `json:"title"`
	Type       string        `json:"type"`
	Status     string        `json:"status"`
	Execution  StepExecution `json:"execution"`
	ApiRetries []ApiRetries  `json:"api_retries"`
	Errors     []error       `json:"errors"`
}

type StepExecution struct {
	MaxRetries    int     `json:"max_retries"`
	RetryInterval float64 `json:"retry_interval"`
	StartupDelay  float64 `json:"startup_delay"`
	FastFail      bool    `json:"fast_fail"`
	Duration      float64 `json:"duration"`
	StartTime     float64 `json:"start_time"`
	EndTime       float64 `json:"end_time"`
}

type ApiRetries struct {
	Request     map[string]string `json:"request"`
	Response    json.RawMessage   `json:"response"`
	RequestUUID string            `json:"request_uuid"`
	RequestTime float64           `json:"request_time"`
}

type Step struct {
	T             *testing.T
	Invoker       func(*Step) (interface{}, error)
	Validators    func(*Step) []TestValidator
	VarsMap       map[string]string
	MaxRetries    int
	RetryInterval time.Duration
	StartupDelay  time.Duration
	FastFail      bool
	StartTime     float64
	EndTime       float64
	Status        string
	Title         string
	Type          string
	Errors        []error
	Owners        []string
	Scenario      *Scenario
	APIRetries    []ApiRetries
}

func (t *Step) NewClient(product string) (interface{}, error) {
	cfg := t.Scenario.Spec.Config
	cred := t.Scenario.Spec.Credential

	client := newServiceClient(product, cfg, cred)
	if client == nil {
		return nil, fmt.Errorf("can not setup client form the %s", product)
	}

	if err := client.AddResponseHandler(t.handleResponse); err != nil {
		return nil, err
	}

	return client, nil
}

// Run will run the test case with retry
func (t *Step) Run() {
	t.StartTime = float64(time.Now().Unix())
	if t.StartupDelay != time.Duration(0) {
		<-time.After(t.StartupDelay)
	}

	defer func() {
		t.EndTime = float64(time.Now().Unix())
	}()

	for i := 0; i < t.MaxRetries+1; i++ {
		t.Errors = []error{}

		resp, err := t.Invoker(t)
		if err != nil {
			if e, ok := err.(uerr.Error); ok && e.Name() == uerr.ErrSendRequest {
				t.Status = "failed"
				t.Errors = append(t.Errors, err)
				t.T.Fatal(err)
				return
			} else if ok && e.Name() == uerr.ErrRetCode {
				// pass
			} else {
				t.Errors = append(t.Errors, err)
				assert.NoError(t.T, err)
				//continue
			}
		}

		for _, validator := range t.Validators(t) {
			if err := validator(resp); err != nil {
				t.Errors = append(t.Errors, err)
			}
		}

		if len(t.Errors) > 0 {
			if i == t.MaxRetries {
				t.Status = "failed"
				t.T.Fatal(err)
				return
			}
			<-time.After(t.RetryInterval)
			continue
		}

		t.Status = "passed"
		t.BindVars(resp)
		return
	}

	return
}

func (t *Step) BindVars(resp interface{}) {
	for k, v := range t.VarsMap {
		t.Scenario.SetVar(k, t.Scenario.Must(utest.GetValue(resp, v)))
	}
}

func (t *Step) Report() StepReport {
	return StepReport{
		Title:  t.Title,
		Type:   t.Type,
		Status: t.Status,
		Execution: StepExecution{
			MaxRetries:    t.MaxRetries,
			RetryInterval: t.RetryInterval.Seconds(),
			StartupDelay:  t.StartupDelay.Seconds(),
			FastFail:      t.FastFail,
			Duration:      t.EndTime - t.StartTime,
			StartTime:     t.StartTime,
			EndTime:       t.EndTime,
		},
		ApiRetries: t.APIRetries,
		Errors:     t.Errors,
	}
}

func (t *Step) init(s *Scenario) {
	t.Status = "skipped"
	t.Type = "api"
	t.Scenario = s
	t.T = s.T
}

func (t *Step) handleResponse(c *ucloud.Client, req request.Common, resp response.Common, retError error) (response.Common, error) {
	if retError != nil {
		if e, ok := retError.(uerr.Error); ok && e.Name() == uerr.ErrRetCode {
		} else {
			return nil, retError
		}
	}
	reqMap, err := request.ToQueryMap(req)
	if err != nil {
		return nil, err
	}
	res, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	t.APIRetries = append(t.APIRetries, ApiRetries{
		Request:     reqMap,
		Response:    res,
		RequestUUID: resp.GetRequestUUID(),
		RequestTime: float64(req.GetRequestTime().Unix()),
	})
	return resp, retError
}
