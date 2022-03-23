package response

type UserLoginResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func ToUserLoginResponse(id string, username string) (userLoginResponse UserLoginResponse) {
	userLoginResponse.Id = id
	userLoginResponse.Username = username
	return userLoginResponse
}
