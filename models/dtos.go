package models

// Expected to get Username & Password on login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Expected to get Username & Password on register request
type RegisterRequest struct {
	Username string `jason:"username"`
	Password string `json:"password"`
}

// Return UserID and Username on success
type LoginSuccessData struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
}

type RegisterSuccessData struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
}

type ErrorDetail struct {
	Message string `json:"message"`
}

type APIResponse struct {
	Success bool         `json:"success"`
	Data    any          `json:"data,omitempty"`
	Error   *ErrorDetail `json:"error,omitempty"`
}

// 'Factory' for success / fail responses
func NewSuccessResponse(data any) APIResponse {
	return APIResponse{
		Success: true,
		Data:    data,
		Error:   nil,
	}
}

func NewErrorResponse(errorMessage string) APIResponse {
	return APIResponse{
		Success: false,
		Data:    nil,
		Error: &ErrorDetail{
			Message: errorMessage,
		},
	}
}
