package response

import "encoding/json"

type GenericResponse interface {
	SetPayload(m map[string]interface{})
	Payload() map[string]interface{}
}
type BaseGenericResponse struct {
	CommonBase

	payload map[string]interface{}
}

func (r *BaseGenericResponse) SetPayload(m map[string]interface{}) {
	r.payload = m
}

func (r BaseGenericResponse) Payload() map[string]interface{} {
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
	body, err := json.Marshal(r.Payload())
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return err
	}
	return nil
}
