package models

import "github.com/google/uuid"

// Session model
type Session struct {
	SessionId string    `json:"session_id" redis:"session_id"`
	UserId    uuid.UUID `json:"user_id" redis:"user_id"`
}

// Base response
type BaseResponse struct {
	Result     string `json:"result"`
	StatusCode int64  `json:"status_code"`
	Message    string `json:"message"`
}
