package response

type LoginResponse struct {
	User      string `json:"user"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}
