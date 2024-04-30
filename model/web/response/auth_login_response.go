package response

type AuthLoginResponse struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expiredAt"`
}
