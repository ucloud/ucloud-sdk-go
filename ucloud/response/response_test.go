package response

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponseAccessor(t *testing.T) {
	resp := CommonBase{
		Action:  "Test",
		RetCode: 42,
		Message: "foo",
	}
	assert.Equal(t, "Test", resp.GetAction())
	assert.Equal(t, 42, resp.GetRetCode())
	assert.Equal(t, "foo", resp.GetMessage())

	resp.SetRequestUUID("foo")
	assert.Equal(t, "foo", resp.GetRequestUUID())

	resp.SetRequest(nil)
	assert.Equal(t, nil, resp.GetRequest())
}

func TestResponseAccessorForGeneric(t *testing.T) {
	resp := BaseGenericResponse{
		CommonBase{
			Action:  "Test",
			RetCode: 1,
			Message: "foo",
		},
		map[string]interface{}{},
	}

	assert.Equal(t, "Test", resp.GetAction())
	assert.Equal(t, 1, resp.GetRetCode())
	assert.Equal(t, "foo", resp.GetMessage())

	assert.NoError(t, resp.SetPayload(map[string]interface{}{
		"Action":  "Test2",
		"RetCode": 2.0,
		"Message": "bar",
	}))

	resp.GetPayload()
	assert.Equal(t, map[string]interface{}{
		"Action":  "Test2",
		"RetCode": 2.0,
		"Message": "bar",
	}, resp.GetPayload())

	assert.Equal(t, "Test2", resp.GetAction())
	assert.Equal(t, 2, resp.GetRetCode())
	assert.Equal(t, "bar", resp.GetMessage())

	assert.Error(t, resp.SetPayload(map[string]interface{}{
		"Message": true,
	}))
	assert.Error(t, resp.SetPayload(map[string]interface{}{
		"Action": 1,
	}))
	assert.Error(t, resp.SetPayload(map[string]interface{}{
		"RetCode": 1,
	}))
}

func TestBaseGenericResponse_Unmarshal(t *testing.T) {
	type testResponse struct {
		CommonBase
		TestIds []string
		TestInt int
	}
	respCreate := &testResponse{}
	type fields struct {
		CommonBase CommonBase
		payload    map[string]interface{}
	}
	type args struct {
		resp interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"OK",
			fields{
				CommonBase: CommonBase{
					Action:  "Test",
					RetCode: 42,
					Message: "foo",
				},
				payload: map[string]interface{}{
					"TestIds": []string{"abc-123"},
					"TestInt": 1,
				},
			},
			args{&testResponse{}},
			false,
		},

		{
			"n_OK",
			fields{
				CommonBase: CommonBase{
					Action:  "Test",
					RetCode: 42,
					Message: "foo",
				},
				payload: map[string]interface{}{
					"TestIds":    []string{"abc-123"},
					"TestString": "test",
				},
			},
			args{&testResponse{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := BaseGenericResponse{
				CommonBase: tt.fields.CommonBase,
				payload:    tt.fields.payload,
			}
			if err := r.Unmarshal(respCreate); err != nil {
				panic(err)
			}
			if err := r.Unmarshal(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
