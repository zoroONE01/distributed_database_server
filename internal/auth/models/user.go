package models

import (
	"time"

	commonModel "distributed_database_server/internal/models"

	"github.com/gofrs/uuid"
)

type RequestList struct {
	commonModel.RequestPaging
	UserName string
}

func (r *RequestList) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"user_name": r.UserName,
		"page":      r.Page,
		"size":      r.Size,
		"sort_by":   r.SortBy,
		"order_by":  r.OrderBy,
	}
}

type UserResponse struct {
	Id          int       `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	UserName    string    `json:"user_name,omitempty"`
	Role        string    `json:"role,omitempty"`
	About       string    `json:"about,omitempty"`
	Avatar      string    `json:"avatar,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Address     string    `json:"address,omitempty"`
	City        string    `json:"city,omitempty"`
	Merchant    string    `json:"merchant,omitempty"`
	Gender      string    `json:"gender,omitempty"`
}

type SaveRequest struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	Birthday  time.Time `json:"birthday"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// User sign in response
type UserWithToken struct {
	User  *UserResponse
	Token string `json:"token"`
}
