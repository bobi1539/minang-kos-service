package response

type AuthLoginResponse struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expiredAt"`
}
