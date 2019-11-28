package driver

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

type SpecificationReport struct {
	Status       string           `json:"status"`
	Execution    Execution        `json:"execution"`
	Scenarios    []ScenarioReport `json:"scenarios"`
	PassedCount  int              `json:"passedCount"`
	FailedCount  int              `json:"failedCount"`
	SkippedCount int              `json:"skippedCount"`
}

type Execution struct {
	MaxRetries    int     `json:"max_retries"`
	RetryInterval float64 `json:"retry_interval"`
	StartupDelay  float64 `json:"startup_delay"`
	FastFail      bool    `json:"fast_fail"`
	Duration      float64 `json:"duration"`
	StartTime     float64 `json:"start_time"`
	EndTime       float64 `json:"end_time"`
}

type Specification struct {
	Scenarios  []*Scenario
	Config     *ucloud.Config
	Credential *auth.Credential
}

func (spec *Specification) AddScenario(s *Scenario) {
	s.Spec = spec
	spec.Scenarios = append(spec.Scenarios, s)
}

func (spec *Specification) Report() SpecificationReport {
	var scenarios []ScenarioReport
	var passedCount, failedCount, skippedCount int
	for _, v := range spec.Scenarios {
		scenarios = append(scenarios, v.Report())
		switch v.status() {
		case "passed":
			passedCount++
		case "failed":
			failedCount++
		case "skipped":
			skippedCount++
		}
	}
	return SpecificationReport{
		Status: spec.status(),
		Execution: Execution{
			Duration:  spec.endTime() - spec.startTime(),
			StartTime: spec.startTime(),
			EndTime:   spec.endTime(),
		},
		Scenarios:    scenarios,
		PassedCount:  passedCount,
		FailedCount:  failedCount,
		SkippedCount: skippedCount,
	}
}

func (spec *Specification) status() string {
	var status []string
	for _, v := range spec.Scenarios {
		switch v.status() {
		case "failed":
			return "failed"
		case "skipped":
			continue
		case "passed":
			status = append(status, "passed")
		}
	}

	if len(status) == 0 {
		return "skipped"
	}

	return "passed"
}

func (spec *Specification) startTime() float64 {
	var t float64
	for _, v := range spec.Scenarios {
		if v.status() != "skipped" {
			if t == 0 {
				t = v.startTime()
			} else if v.startTime() < t {
				t = v.startTime()
			}
		}
	}
	return t
}

func (spec *Specification) endTime() float64 {
	var t float64
	for _, v := range spec.Scenarios {
		if v.status() != "skipped" && v.endTime() > t {
			t = v.endTime()
		}
	}

	return t
}
