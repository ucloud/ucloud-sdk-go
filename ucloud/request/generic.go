package request

type GenericRequest interface {
	Common

	SetPayload(m map[string]interface{})
	GetPayload() map[string]interface{}
}

type BaseGenericRequest struct {
	CommonBase

	payload map[string]interface{}
}

func (r *BaseGenericRequest) SetPayload(m map[string]interface{}) {
	r.payload = m
}

func (r BaseGenericRequest) GetPayload() map[string]interface{} {
	m := make(map[string]interface{})
	if len(r.CommonBase.GetRegion()) != 0 {
		m["Region"] = r.CommonBase.GetRegion()
	}

	if len(r.CommonBase.GetZone()) != 0 {
		m["Zone"] = r.CommonBase.GetZone()
	}

	if len(r.CommonBase.GetAction()) != 0 {
		m["Action"] = r.CommonBase.GetAction()
	}

	if len(r.CommonBase.GetProjectId()) != 0 {
		m["ProjectId"] = r.CommonBase.GetProjectId()
	}

	for k, v := range r.payload {
		m[k] = v
	}

	return m
}

func (r *BaseGenericRequest) GetAction() string {
	if r.payload["Action"] != nil {
		return r.payload["Action"].(string)
	}

	return r.CommonBase.GetAction()
}

func (r *BaseGenericRequest) GetRegion() string {
	if r.payload["Region"] != nil {
		return r.payload["Region"].(string)
	}

	return r.CommonBase.GetRegion()
}

func (r *BaseGenericRequest) GetZone() string {
	if r.payload["Zone"] != nil {
		return r.payload["Zone"].(string)
	}

	return r.CommonBase.GetZone()
}

func (r *BaseGenericRequest) GetProjectId() string {
	if r.payload["ProjectId"] != nil {
		return r.payload["ProjectId"].(string)
	}

	return r.CommonBase.GetProjectId()
}
