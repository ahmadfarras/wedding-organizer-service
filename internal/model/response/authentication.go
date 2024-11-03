package response

type LoginResponse struct {
	Token string `json:"token"`
}

func BuildLoginResponse(token string) *LoginResponse {
	return &LoginResponse{
		Token: token,
	}
}
