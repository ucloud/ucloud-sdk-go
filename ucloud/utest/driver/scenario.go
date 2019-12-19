package driver

import (
	"encoding/json"
	"testing"
)

type executionErrors []error

func (err executionErrors) MarshalJSON() ([]byte, error) {
	var errString []string
	for _, v := range err {
		errString = append(errString, v.Error())
	}
	return json.Marshal(errString)
}

type ScenarioReport struct {
	Title        string            `json:"title"`
	Id           string            `json:"id"`
	Steps        []StepReport      `json:"steps"`
	Status       string            `json:"status"`
	Execution    ScenarioExecution `json:"execution"`
	Owners       []string          `json:"owners"`
	PassedCount  int               `json:"passedCount"`
	FailedCount  int               `json:"failedCount"`
	SkippedCount int               `json:"skippedCount"`
	Errors       executionErrors   `json:"errors,omitempty"`
}

type ScenarioExecution struct {
	Duration  float64 `json:"duration"`
	StartTime float64 `json:"start_time"`
	EndTime   float64 `json:"end_time"`
}

type Scenario struct {
	PreCheck func()
	Id       string
	Title    string
	Steps    []*Step
	Spec     *Specification
	Owners   []string
	Vars     func(*Scenario) map[string]interface{}
	Errors   []error

	vars map[string]interface{}
}

// SetVar will set the variable of Scenario
func (scenario *Scenario) SetVar(name string, value interface{}) {
	scenario.vars[name] = value
}

// GetVar will return the variable of Scenario
func (scenario *Scenario) GetVar(name string) interface{} {
	if v, ok := scenario.vars[name]; ok {
		return v
	}
	return ""
}

// Must will check error is nil and return the value
func (scenario *Scenario) Must(v interface{}, err error) interface{} {
	if err != nil {
		scenario.Errors = append(scenario.Errors, err)
	}
	return v
}

func (scenario *Scenario) init() {
	scenario.vars = map[string]interface{}{}
}

// Run will run the scenario test case
func (scenario *Scenario) Run(t *testing.T) {
	scenario.init()
	for k, v := range scenario.Vars(scenario) {
		scenario.SetVar(k, v)
	}

	for i := 0; i < len(scenario.Steps); i++ {
		step := scenario.Steps[i]
		step.init()
		step.Scenario = scenario
		step.T = t
		step.id = i + 1
		step.run()
		if len(step.Errors) > 0 && step.FastFail {
			t.Fatal(step.Errors)
			return
		}
	}
}

func (scenario *Scenario) Report() ScenarioReport {
	var steps []StepReport
	var passedCount, failedCount, skippedCount int
	for _, v := range scenario.Steps {
		steps = append(steps, v.Report())
		switch v.Status {
		case "passed":
			passedCount++
		case "failed":
			failedCount++
		case "skipped":
			skippedCount++
		}
	}
	return ScenarioReport{
		Title:  scenario.Title,
		Id:     scenario.Id,
		Status: scenario.status(),
		Execution: ScenarioExecution{
			Duration:  scenario.endTime() - scenario.startTime(),
			StartTime: scenario.startTime(),
			EndTime:   scenario.endTime(),
		},
		Owners:       scenario.Owners,
		PassedCount:  passedCount,
		FailedCount:  failedCount,
		SkippedCount: skippedCount,
		Steps:        steps,
		Errors:       scenario.Errors,
	}
}

func (scenario *Scenario) startTime() float64 {
	var t float64
	for _, v := range scenario.Steps {
		if v.Status != "skipped" && v.StartTime != 0 {
			if t == 0 {
				t = v.StartTime
			} else if v.StartTime < t {
				t = v.StartTime
			}
		}
	}
	return t
}

func (scenario *Scenario) endTime() float64 {
	var t float64
	for _, v := range scenario.Steps {
		if v.Status != "skipped" && v.EndTime > t {
			t = v.EndTime
		}
	}
	return t
}

func (scenario *Scenario) status() string {
	var status []string
	for _, v := range scenario.Steps {
		switch v.Status {
		case "failed":
			return "failed"
		case "skipped":
			continue
		case "passed":
			status = append(status, v.Status)
		}
	}

	if len(status) == 0 {
		return "skipped"
	}
	return "passed"
}
