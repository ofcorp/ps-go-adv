package verify

type SendRequest struct {
	Email    string `json:"email" validate:"required,email"`
}

type SendResponse struct {
	Hash string `json:"hash"`
}

type VerifyResponse struct {
	Result bool `json:"result"`
}