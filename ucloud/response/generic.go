package response

import "encoding/json"

type GenericResponse interface {
	Common
	SetPayload(m map[string]interface{})
	GetPayload() map[string]interface{}
	Unmarshal(interface{}) error
}
type BaseGenericResponse struct {
	CommonBase

	payload map[string]interface{}
}

func (r *BaseGenericResponse) SetPayload(m map[string]interface{}) {
	r.payload = m
}

func (r BaseGenericResponse) GetPayload() map[string]interface{} {
	m := make(map[string]interface{})

	m["RetCode"] = r.GetRetCode()
	m["Action"] = r.GetAction()
	m["Message"] = r.GetMessage()
	m["Action"] = r.GetAction()

	for k, v := range r.payload {
		m[k] = v
	}

	return m
}

func (r BaseGenericResponse) Unmarshal(resp interface{}) error {
	body, err := json.Marshal(r.GetPayload())
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return err
	}
	return nil
}
