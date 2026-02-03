package auth

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
}

type LoginResponse struct {
	SessionID string `json:"session_id"`
}

type VerifyRequest struct {
	SessionID    string `json:"session_id" validate:"required"`
	VerificationCode string `json:"code" validate:"required"`
}

type VerifyResponse struct {
	Token string `json:"token"`
}
