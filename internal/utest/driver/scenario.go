package driver

import (
	"github.com/stretchr/testify/assert"
	"github.com/ucloud/ucloud-sdk-go/internal/utest"
	"testing"
)

type ScenarioReport struct {
	Title        string       `json:"title"`
	Steps        []StepReport `json:"steps"`
	Status       string       `json:"status"`
	Execution    Execution    `json:"execution"`
	Owners       []string     `json:"owners"`
	PassedCount  int          `json:"passedCount"`
	FailedCount  int          `json:"failedCount"`
	SkippedCount int          `json:"skippedCount"`
}

type Scenario struct {
	T           *testing.T
	Id          int
	Title       string
	Steps       []*Step
	Spec        *Specification
	Owners      []string
	Vars        map[string]interface{}
	Comparators *utest.Comparators
}

// SetVar will set the variable of Scenarios
func (s *Scenario) SetVar(name string, value interface{}) {
	s.Vars[name] = value
}

// GetVar will return the variable of Scenarios
func (s *Scenario) GetVar(name string) interface{} {
	if v, ok := s.Vars[name]; ok {
		return v
	}
	return ""
}

// MustString will check error is nil and return the string value
func (s *Scenario) MustString(v string, err error) string {
	assert.NoError(s.T, err)
	return v
}

// Must will check error is nil and return the value
func (s *Scenario) Must(v interface{}, err error) interface{} {
	assert.NoError(s.T, err)
	return v
}

// NoError will check error is nil
func (s *Scenario) NoError(err error) {
	assert.NoError(s.T, err)
}

func (s *Scenario) Run() {
	for i := 0; i < len(s.Steps); i++ {
		step := s.Steps[i]
		step.Run()
		if len(step.Errors) > 0 && step.FastFail {
			return
		}
	}
}

func (s *Scenario) AddStep(step *Step) {
	step.init(s)
	s.Steps = append(s.Steps, step)
}

func (s *Scenario) Report() ScenarioReport {
	var steps []StepReport
	var passedCount, failedCount, skippedCount int
	for _, v := range s.Steps {
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
		Title:  s.Title,
		Steps:  steps,
		Status: s.status(),
		Execution: Execution{
			Duration:  s.endTime() - s.startTime(),
			StartTime: s.startTime(),
			EndTime:   s.endTime(),
		},
		Owners:       s.Owners,
		PassedCount:  passedCount,
		FailedCount:  failedCount,
		SkippedCount: skippedCount,
	}
}

func (s *Scenario) startTime() float64 {
	var t float64
	for _, v := range s.Steps {
		if v.Status != "skipped" {
			if t == 0 {
				t = v.StartTime
			} else if v.StartTime < t {
				t = v.StartTime
			}
		}
	}
	return t
}

func (s *Scenario) endTime() float64 {
	var t float64
	for _, v := range s.Steps {
		if v.Status != "skipped" && v.EndTime > t {
			t = v.EndTime
		}
	}
	return t
}

func (s *Scenario) status() string {
	var status []string
	for _, v := range s.Steps {
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
