package waiter

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestWaitForState(t *testing.T) {
	tests := []struct {
		name    string
		conf    *StateWaiter
		want    interface{}
		wantErr bool
	}{
		{
			"ok",
			&StateWaiter{
				Pending: []string{"pending"},
				Target:  []string{"ok"},
				Refresh: func() (interface{}, string, error) {
					return true, "ok", nil
				},
				Timeout: 2 * time.Second,
			},
			true,
			false,
		},
		{
			"timeout",
			&StateWaiter{
				Pending: []string{"pending"},
				Target:  []string{"ok"},
				Refresh: func() (interface{}, string, error) {
					return true, "pending", nil
				},
				Timeout:    2 * time.Second,
				MinTimeout: 1 * time.Second,
			},
			nil,
			true,
		},
		{
			"confTimeout",
			&StateWaiter{
				Pending: []string{"pending"},
				Target:  []string{"ok"},
				Refresh: func() (interface{}, string, error) {
					return true, "ok", nil
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.conf.Wait()
			if (err != nil) != tt.wantErr {
				t.Errorf("StateWaiter.Wait() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StateWaiter.Wait() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorToString(t *testing.T) {
	err := TimeoutError{
		LastError:      errTimeoutConf,
		LastState:      "pending",
		Timeout:        1 * time.Second,
		ExpectedStates: []string{"ok"},
	}
	assert.NotZero(t, err.Error())
}
