package response

type AuthResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
}

type CommonA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
}
