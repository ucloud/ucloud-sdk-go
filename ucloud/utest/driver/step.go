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
	Owners        []string
	Scenario      *Scenario

	Invoker    func(*Step) (interface{}, error)
	Validators func(*Step) []TestValidator

	errors     []error
	apiRetries []ApiRetries

	startTime float64
	endTime   float64
	status    string
	id        int
}

// LoadFixture is a function for load fixture by the name from map fixture function
func (step *Step) LoadFixture(name string) (interface{}, error) {
	if step.Scenario.Spec.fixtures[name] != nil {
		return step.Scenario.Spec.fixtures[name](step)
	}
	return nil, fmt.Errorf("can not load fixture by the %s", name)
}

// SetupClientFixture is a help function for setup client fixture
func SetupClientFixture(client ucloud.ServiceClient) FixtureFunc {
	return func(step *Step) (i interface{}, e error) {
		if err := client.AddResponseHandler(step.handleResponse); err != nil {
			return nil, err
		}

		return client, nil
	}
}

// Must will check error is nil and return the value
func (step *Step) Must(v interface{}, err error) interface{} {
	if err != nil {
		step.AppendError(err)
	}
	return v
}

func (step *Step) AppendApiRetries(apiRetry ApiRetries) {
	step.apiRetries = append(step.apiRetries, apiRetry)
}

func (step *Step) AppendError(err error) {
	step.errors = append(step.errors, fmt.Errorf("step %02d Failed, %s", step.id, err))
}

// Run will run the step test case with retry
func (step *Step) run() {
	step.startTime = float64(time.Now().Unix())
	if step.StartupDelay != time.Duration(0) {
		<-time.After(step.StartupDelay)
	}

	defer func() {
		step.endTime = float64(time.Now().Unix())
	}()

	for i := 0; i < step.MaxRetries+1; i++ {
		step.errors = []error{}

		resp, err := step.Invoker(step)
		if err != nil {
			if e, ok := err.(uerr.Error); ok && e.Name() == uerr.ErrSendRequest {
				step.status = "failed"
				step.AppendError(err)
				assert.NoError(step.T, err)
				return
			} else if ok && e.Name() == uerr.ErrRetCode {
				// pass
			} else {
				step.AppendError(err)
				// continue
			}
		}

		for _, validator := range step.Validators(step) {
			if err := validator(resp); err != nil {
				step.AppendError(err)
			}
		}

		if len(step.errors) > 0 {
			if i == step.MaxRetries {
				step.status = "failed"
				return
			}
			<-time.After(step.RetryInterval)
			continue
		}

		step.status = "passed"
		return
	}

	return
}

func (step *Step) Report() StepReport {
	return StepReport{
		Title:  step.Title,
		Status: step.status,
		Execution: StepExecution{
			MaxRetries:    step.MaxRetries,
			RetryInterval: step.RetryInterval.Seconds(),
			StartupDelay:  step.StartupDelay.Seconds(),
			FastFail:      step.FastFail,
			Duration:      step.endTime - step.startTime,
			StartTime:     step.startTime,
			EndTime:       step.endTime,
		},
		ApiRetries: step.apiRetries,
		Errors:     step.errors,
	}
}

func (step *Step) init() {
	step.status = "skipped"
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

	step.apiRetries = append(step.apiRetries, ApiRetries{
		Request:     reqMap,
		Response:    res,
		RequestUUID: resp.GetRequestUUID(),
		RequestTime: float64(req.GetRequestTime().Unix()),
	})
	return resp, retError
}
