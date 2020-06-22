package models

import (
	"log"
	_ "os"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	u "bank-example/internal/utils"
)

const (
	DECLINED int = iota
	APPROVED
)

type (
	Users struct {
		ID        int64  `json:"-"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password,omitempty"`
	}
	UserBalance struct {
		Balance int64 `json:"balanceInCents"`
	}
	Auth struct {
		ID            int64
		Email         string
		Authenticated bool
	}
)

//Validate incoming user details...
func (user *Users) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	if len(user.FirstName) <= 0 {
		return u.Message(false, "First Name is required"), false
	}

	if len(user.LastName) <= 0 {
		return u.Message(false, "Last Name is required"), false
	}

	//Email must be unique
	temp := &Users{}

	//check for errors and duplicate emails
	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

//Validate incoming user details...
func (user *Users) ValidateUpdate() (map[string]interface{}, bool) {
	result := make(map[string]interface{})
	if user.Email != "" {
		if !strings.Contains(user.Email, "@") {
			return u.Message(false, "Email address is incorrect"), false
		}
		//Email must be unique
		temp := &Users{}

		//check for errors and duplicate emails
		err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
		if err != nil && err == gorm.ErrRecordNotFound {
			return u.Message(false, "Connection error. Please retry"), false
		}
		if temp.Email == "" {
			return u.Message(false, "Invalid Request"), false
		}
		result["email"] = user.Email
	}

	if user.Password != "" {
		if len(user.Password) < 6 {
			return u.Message(false, "Password length should be greater than 6"), false
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = ""
		result["password"] = string(hashedPassword)
	}

	if user.FirstName != "" {
		if len(user.FirstName) <= 0 {
			return u.Message(false, "First Name is required"), false
		}
		result["first_name"] = user.FirstName
	}

	if user.LastName != "" {
		if len(user.LastName) <= 0 {
			return u.Message(false, "Last Name is required"), false
		}
		result["last_name"] = user.LastName
	}

	return result, true
}

func (user *Users) Create() map[string]interface{} {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error.")
	}

	user.Password = ""

	response := u.Message(true, "User has been created")
	response["user"] = user
	return response
}

func (user *Users) Update(id int64) map[string]interface{} {
	resp, ok := user.ValidateUpdate()
	if !ok {
		return resp
	}
	if len(resp) <= 0 {
		return u.Message(false, "Invalid Request")
	}
	GetDB().Table("users").Where("id = ?", id).Updates(resp)

	response := u.Message(true, "User has been updated")
	return response
}

func AuthUser(email, password string) *Auth {
	auth := &Auth{}
	err := GetDB().Raw(`
		SELECT id AS ID, email AS Email, password=crypt(?, password) AS Authenticated FROM users
		WHERE email= ?
	`, password, email).Scan(auth).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		log.Print("Connection error. Please retry")
		return nil
	}

	return auth
}

func (user *Users) Get(id int64) map[string]interface{} {
	GetDB().Table("users").Where("id = ?", id).First(user)
	if user.Email == "" {
		return nil
	}

	user.Password = ""
	response := u.Message(true, "success")
	response["user"] = user
	return response
}

func (user *Users) GetBalance(userID int64) map[string]interface{} {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("Fetch took %s", elapsed)
	}()
	result := UserBalance{}
	err := GetDB().Raw(`
		SELECT SUM(transactions.amountInCents) AS balance FROM transactions
		INNER JOIN users ON users.id = transactions.user_id
		WHERE users.id = ?
	`, userID).Scan(&result).Error
	if err != nil {
		log.Printf("Error: %s", err)
		return nil
	}
	resp := u.Message(true, "success")
	resp["balance"] = result
	return resp
}

func (user *Users) AuthorizeTransaction(transaction *Transaction) map[string]interface{} {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("Fetch took %s", elapsed)
	}()
	result := UserBalance{}
	err := GetDB().Raw(`
		SELECT SUM(transactions.amountInCents) AS balance FROM transactions
		INNER JOIN users ON users.id = transactions.user_id
		WHERE users.id = ?
	`, user.ID).Scan(&result).Error
	if err != nil {
		log.Printf("Error: %s", err)
		return nil
	}
	resp := u.Message(false, "not enough balance")
	resp["transaction"] = DECLINED
	if (result.Balance + transaction.AmountInCents) > 0 {
		resp = u.Message(true, "success")
		resp["balance"] = APPROVED
	}
	return resp
}
