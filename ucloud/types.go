package ucloud

type CommonRequest struct {
	Action    string
	PublicKey string
	ProjectId string
	Signature string
}

type CommonResponse struct {
	Action  string
	RetCode int
}
