package waiter

import (
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
				Timeout: 2 * time.Second,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.conf.Wait()
			if (err != nil) != tt.wantErr {
				t.Errorf("StateChangeConf.WaitForState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StateChangeConf.WaitForState() = %v, want %v", got, tt.want)
			}
		})
	}
}
