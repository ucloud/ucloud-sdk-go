package utest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/helpers/mock"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestDriver(t *testing.T) {
	type input struct {
		Id          string
		Title       string
		Owners      []string
		Steps       []*driver.Step
		Mock        map[string]interface{}
		InitialVars map[string]interface{}
	}

	type golden struct {
		stepFailedCount  int
		stepPassedCount  int
		stepSkippedCount int
	}

	tests := []struct {
		name   string
		input  input
		golden golden
	}{
		{
			"ok",
			input{
				Steps: []*driver.Step{
					{
						Invoker: func(step *driver.Step) (interface{}, error) {
							return response.CommonBase{RetCode: 0}, nil
						},
						Validators: func(step *driver.Step) []driver.TestValidator {
							return []driver.TestValidator{
								validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
							}
						},
						Title: "step1",
					},
				},
				Id:     "10",
				Title:  "test_ok",
				Owners: []string{"foo"},
				InitialVars: map[string]interface{}{
					"foo": 0,
					"bar": "1",
				},
			},
			golden{
				stepFailedCount:  0,
				stepPassedCount:  1,
				stepSkippedCount: 0,
			},
		},
		{
			"failed",
			input{
				Steps: []*driver.Step{
					{
						Invoker: func(step *driver.Step) (interface{}, error) {
							return response.CommonBase{RetCode: 0}, fmt.Errorf("foo")
						},
						Validators: func(step *driver.Step) []driver.TestValidator {
							return []driver.TestValidator{
								validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
							}
						},
					},
				},
			},
			golden{
				stepFailedCount:  1,
				stepPassedCount:  0,
				stepSkippedCount: 0,
			},
		},
		{
			"customValidator",
			input{
				Steps: []*driver.Step{
					{
						Invoker: func(step *driver.Step) (interface{}, error) {
							return response.CommonBase{RetCode: 0}, nil
						},
						Validators: func(step *driver.Step) []driver.TestValidator {
							return []driver.TestValidator{
								func(resp interface{}) error {
									if !reflect.DeepEqual(step.Must(utils.GetValue(resp, "RetCode")), 0) {
										return fmt.Errorf("foo")
									}
									return nil
								},
							}
						},
					},
				},
			},
			golden{
				stepFailedCount:  0,
				stepPassedCount:  1,
				stepSkippedCount: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var spec = &driver.Specification{}

			scenario := &driver.Scenario{
				Id:     tt.input.Id,
				Title:  tt.input.Title,
				Steps:  tt.input.Steps,
				Owners: tt.input.Owners,
				Vars: func(scenario *driver.Scenario) map[string]interface{} {
					return tt.input.InitialVars
				},
			}

			spec.AddScenario(scenario)
			scenario.Run(t)
			got := spec.Report()

			if tt.golden.stepFailedCount != 0 {
				assert.Equal(t, "failed", got.Status)
			}
			assert.Equal(t, tt.golden.stepFailedCount, got.Scenarios[0].FailedCount)
			assert.Equal(t, tt.golden.stepSkippedCount, got.Scenarios[0].SkippedCount)
			assert.Equal(t, tt.golden.stepPassedCount, got.Scenarios[0].PassedCount)

			assert.Equal(t, tt.input.Title, got.Scenarios[0].Title)
			assert.Equal(t, tt.input.Id, got.Scenarios[0].Id)
			assert.Equal(t, tt.input.Owners, got.Scenarios[0].Owners)
			assert.Equal(t, tt.input.Steps[0].Title, got.Scenarios[0].Steps[0].Title)
			// TODO: check information in report is equal to input step information
		})
	}
}

func TestDriverWithClient(t *testing.T) {
	type input struct {
		mockData map[string]interface{}
		mockErr  error
	}

	type golden struct {
		status  string
		retries int
	}

	tests := []struct {
		name   string
		input  input
		golden golden
	}{
		{
			"ok",
			input{
				mockData: map[string]interface{}{"RetCode": 0},
				mockErr:  nil,
			},
			golden{status: "passed", retries: 0},
		},
		{
			"mockSendRequestErr",
			input{
				mockData: map[string]interface{}{},
				mockErr:  fmt.Errorf("foo"),
			},
			golden{status: "failed", retries: 0},
		},
		{
			"mockRetCodeErr",
			input{
				mockData: map[string]interface{}{"RetCode": 10},
				mockErr:  nil,
			},
			golden{status: "passed", retries: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//	t.Parallel()

			var spec = &driver.Specification{}
			spec.AddFixture("mock", func(step *driver.Step) (i interface{}, e error) {
				return newMockedClient(tt.input.mockData, tt.input.mockErr)
			})

			scenario := &driver.Scenario{
				Steps: []*driver.Step{
					{
						MaxRetries: 3,
						Invoker: func(step *driver.Step) (i interface{}, e error) {
							v := step.Must(step.LoadFixture("mock"))
							client := v.(*ucloud.Client)
							req := request.CommonBase{}
							resp := response.CommonBase{}

							err := client.InvokeAction("Test", &req, &resp)
							return resp, err
						},
					},
				},
			}
			spec.AddScenario(scenario)
			scenario.Run(t)

			got := spec.Report()
			if tt.golden.status == "failed" {
				assert.NotZero(t, got.Scenarios[0].Steps[0].Errors)
			}
			assert.Equal(t, tt.golden.status, got.Scenarios[0].Status)
			assert.Equal(t, tt.golden.retries, len(got.Scenarios[0].Steps[0].Retries.Rows))
		})
	}
}

func newMockedClient(mockData map[string]interface{}, mockError error) (*ucloud.Client, error) {
	cfg := ucloud.NewConfig()
	credential := auth.NewCredential()
	client := ucloud.NewClient(&cfg, &credential)
	httpClient := mock.NewHttpClient()

	err := httpClient.MockData(func(requests mock.Request, responses mock.Response) error {
		if mockError != nil {
			return mockError
		}
		for k, v := range mockData {
			responses[k] = v
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if err := client.SetHttpClient(httpClient); err != nil {
		return nil, err
	}
	return client, nil
}

func newStepInvoker() {
	return
}
