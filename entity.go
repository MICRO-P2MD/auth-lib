package auth

type GetCurrentUserResponse struct {
	Status string          `json:"status"`
	Code   int32           `json:"code"`
	Data   CurrentUserData `json:"data"`
}

type CurrentUserData struct {
	User User `json:"user"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type ValidateSecretKeyResponse struct {
	Status string    `json:"status"`
	Code   int32     `json:"code"`
	Data   SecretKey `json:"data"`
}

type SecretKey struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
