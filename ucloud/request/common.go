package request

type GenericRequest interface {
	SetPayload(m map[string]interface{})
	Payload() map[string]interface{}
}

type BaseGenericRequest struct {
	CommonBase

	payload map[string]interface{}
}

func (r *BaseGenericRequest) SetPayload(m map[string]interface{}) {
	r.payload = m
}

func (r BaseGenericRequest) Payload() map[string]interface{} {
	m := make(map[string]interface{})
	if len(r.GetRegion()) != 0 {
		m["Region"] = r.GetRegion()
	}

	if len(r.GetZone()) != 0 {
		m["Zone"] = r.GetZone()
	}

	if len(r.GetAction()) != 0 {
		m["Action"] = r.GetAction()
	}

	if len(r.GetProjectId()) != 0 {
		m["ProjectId"] = r.GetProjectId()
	}

	for k, v := range r.payload {
		m[k] = v
	}

	return m
}
