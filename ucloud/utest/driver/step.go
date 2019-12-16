package driver

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	uerr "github.com/ucloud/ucloud-sdk-go/ucloud/error"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

type StepReport struct {
	Title      string          `json:"title"`
	Type       string          `json:"type"`
	Status     string          `json:"status"`
	Execution  StepExecution   `json:"execution"`
	ApiRetries []ApiRetries    `json:"api_retries"`
	Errors     executionErrors `json:"errors,omitempty"`
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

type TestValidator func(interface{}) error
type Step struct {
	T *testing.T

	MaxRetries    int
	RetryInterval time.Duration
	StartupDelay  time.Duration
	FastFail      bool
	Title         string
	Type          string
	Owners        []string

	Invoker    func(*Step) (interface{}, error)
	Validators func(*Step) []TestValidator

	Errors     []error
	StartTime  float64
	EndTime    float64
	Status     string
	Scenario   *Scenario
	APIRetries []ApiRetries

	id int
}

// NewClient will build a service Client and add response handler
func (step *Step) NewClient(product string) (interface{}, error) {
	cfg := step.Scenario.Spec.Config
	cred := step.Scenario.Spec.Credential

	client := newServiceClient(product, cfg, cred)
	if client == nil {
		return nil, fmt.Errorf("can not setup client form the %s", product)
	}

	if err := client.AddResponseHandler(step.handleResponse); err != nil {
		return nil, err
	}

	return client, nil
}

// Must will check error is nil and return the value
func (step *Step) Must(v interface{}, err error) interface{} {
	if err != nil {
		step.appendError(err)
	}
	return v
}

func (step *Step) appendError(err error) {
	step.Errors = append(step.Errors, fmt.Errorf("step %02d Failed, %s", step.id, err))
}

// Run will run the step test case with retry
func (step *Step) run() {
	step.StartTime = float64(time.Now().Unix())
	if step.StartupDelay != time.Duration(0) {
		<-time.After(step.StartupDelay)
	}

	defer func() {
		step.EndTime = float64(time.Now().Unix())
	}()

	for i := 0; i < step.MaxRetries+1; i++ {
		step.Errors = []error{}

		resp, err := step.Invoker(step)
		if err != nil {
			if e, ok := err.(uerr.Error); ok && e.Name() == uerr.ErrSendRequest {
				step.Status = "failed"
				step.appendError(err)
				assert.NoError(step.T, err)
				return
			} else if ok && e.Name() == uerr.ErrRetCode {
				// pass
			} else {
				step.appendError(err)
				// continue
			}
		}

		for _, validator := range step.Validators(step) {
			if err := validator(resp); err != nil {
				step.appendError(err)
			}
		}

		if len(step.Errors) > 0 {
			if i == step.MaxRetries {
				step.Status = "failed"
				return
			}
			<-time.After(step.RetryInterval)
			continue
		}

		step.Status = "passed"
		return
	}

	return
}

func (step *Step) Report() StepReport {
	return StepReport{
		Title:  step.Title,
		Type:   step.Type,
		Status: step.Status,
		Execution: StepExecution{
			MaxRetries:    step.MaxRetries,
			RetryInterval: step.RetryInterval.Seconds(),
			StartupDelay:  step.StartupDelay.Seconds(),
			FastFail:      step.FastFail,
			Duration:      step.EndTime - step.StartTime,
			StartTime:     step.StartTime,
			EndTime:       step.EndTime,
		},
		ApiRetries: step.APIRetries,
		Errors:     step.Errors,
	}
}

func (step *Step) init() {
	step.Status = "skipped"
	step.Type = "api"
}

func (step *Step) handleResponse(c *ucloud.Client, req request.Common, resp response.Common, retError error) (response.Common, error) {
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

	step.APIRetries = append(step.APIRetries, ApiRetries{
		Request:     reqMap,
		Response:    res,
		RequestUUID: resp.GetRequestUUID(),
		RequestTime: float64(req.GetRequestTime().Unix()),
	})
	return resp, retError
}
