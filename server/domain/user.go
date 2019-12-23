package domain

import (
	"errors"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

type User struct {
	Id        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserService interface {

	/*
	 * Create user
	 */
	Save(user *User) (*User, error)

	/*
	 * Get user
	 */
	Get(id uint64) *User

	FindUserByEmail(email string) *User
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Id = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			return errors.New("El nombre de usuario es necesario")
		}
		if u.Password == "" {
			return errors.New("La clave es necesaria")
		}
		if u.Email == "" {
			return errors.New("La dirección de correo es necesaria")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("La dirección de correo es inválida")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("La clave es necesaria")
		}
		if u.Username == "" {
			return errors.New("El nombre de usuario es necesario")
		}
		return nil

	default:
		if u.Username == "" {
			return errors.New("El nombre de usuario es necesario")
		}
		if u.Password == "" {
			return errors.New("La clave es necesaria")
		}
		if u.Email == "" {
			return errors.New("La dirección de correo es necesaria")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("La dirección de correo es inválida")
		}
		return nil
	}
}
