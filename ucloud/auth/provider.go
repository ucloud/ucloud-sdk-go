package auth

type CredentialProvider interface {
	Retrieve() (Credential, error)
}
