package utils

import (
	"testing"
	"time"
)

func TestFuncWaiter_WaitForCompletion(t *testing.T) {
	tests := []struct {
		name    string
		w       *FuncWaiter
		wantErr bool
	}{
		{
			"OK",
			&FuncWaiter{
				Checker: func() func() (bool, error) {
					i := 0
					return func() (bool, error) {
						i++
						return i == 1, nil
					}
				}(),
				MaxAttempts: 2,
				Interval:    time.Second * 1,
			},
			false,
		},
		{
			"MaxattemptsReached",
			&FuncWaiter{
				Checker: func() func() (bool, error) {
					i := 0
					return func() (bool, error) {
						i++
						return i == 2, nil
					}
				}(),
				MaxAttempts: 1,
				Interval:    time.Second * 1,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.WaitForCompletion(); (err != nil) != tt.wantErr {
				t.Errorf("FuncWaiter.WaitForCompletion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncWaiter_Cancel(t *testing.T) {
	i := 0
	waiter := &FuncWaiter{
		Checker: func() (bool, error) {
			i++
			return i == 3, nil
		},
		MaxAttempts: 3,
		Interval:    time.Second * 1,
	}

	go waiter.WaitForCompletion()
	time.Sleep(time.Millisecond * 500)

	if i > 2 {
		t.Errorf("FuncWaiter.Cancel() failed")
	}
}
