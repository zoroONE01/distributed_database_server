package entity

import (
	"distributed_database_server/internal/auth/models"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Id          int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	UserId      uuid.UUID `gorm:"column:user_id" json:"user_id" redis:"user_id" validate:"omitempty"`
	FirstName   string    `gorm:"column:first_name;default:(-)" json:"first_name" redis:"first_name" validate:"required,lte=30"`
	LastName    string    `gorm:"column:last_name;default:(-)" json:"last_name" redis:"last_name" validate:"required,lte=30"`
	Email       string    `gorm:"column:email" json:"email,omitempty" redis:"email" validate:"omitempty,lte=60,email"`
	Password    string    `gorm:"column:password" json:"password,omitempty" redis:"password" validate:"omitempty,required,gte=6"`
	Role        string    `gorm:"column:role;default:(-)" json:"role,omitempty" redis:"role" validate:"omitempty,lte=10"`
	About       string    `gorm:"column:about;default:(-)" json:"about,omitempty" redis:"about" validate:"omitempty,lte=1024"`
	Avatar      string    `gorm:"column:avatar;default:(-)" json:"avatar,omitempty" redis:"avatar" validate:"omitempty,lte=512,url"`
	PhoneNumber string    `gorm:"column:phone_number;default:(-)" json:"phone_number,omitempty" redis:"phone_number" validate:"omitempty,lte=20"`
	Address     string    `gorm:"column:address;default:(-)" json:"address,omitempty" redis:"address" validate:"omitempty,lte=250"`
	City        string    `gorm:"column:city;default:(-)" json:"city,omitempty" redis:"city" validate:"omitempty,lte=24"`
	Country     string    `gorm:"column:country;default:(-)" json:"country,omitempty" redis:"country" validate:"omitempty,lte=24"`
	Gender      string    `gorm:"column:gender" json:"gender,omitempty" redis:"gender" validate:"omitempty,lte=10"`
	Birthday    time.Time `gorm:"column:birthday;default:(-)" json:"birthday,omitempty" redis:"birthday" validate:"omitempty,lte=10"`
	CreatedAt   time.Time `gorm:"autoCreatetime" json:"created_at,omitempty" redis:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at,omitempty" redis:"updated_at"`
	LoginDate   time.Time `gorm:"column:login_date;default:(-)" json:"login_date,omitempty" redis:"login_date" validate:"omitempty"`
}

func (u *User) TableName() string {
	return "users"
}

// Hash user password with bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Sanitize user password
func (u *User) SanitizePassword() {
	u.Password = ""
}

// Generate new user id
func (u *User) newUUID() {
	u.UserId = uuid.New()
}

func (u *User) Export() *models.UserResponse {
	obj := &models.UserResponse{}
	copier.Copy(obj, u) //nolint
	return obj
}

func (u *User) ParseFromSaveRequest(req *models.SaveRequest) {
	copier.Copy(u, req) //nolint
	u.newUUID()
	u.HashPassword()
}
