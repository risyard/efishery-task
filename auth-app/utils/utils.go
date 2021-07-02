package utils

type User struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`
	Password   string `json:"password"`
	Timestampz string `json:"timestamp"`
}

type UserResponse struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type SuccessResponse struct {
	Status int         `json:"status_code"`
	Data   interface{} `json:"data"`
}

type BadResponse struct {
	Status  int    `json:"status_code"`
	Message string `json:"message"`
}
