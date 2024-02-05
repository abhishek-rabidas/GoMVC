package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"log"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func UnMarshalUser(request io.ReadCloser) *User {
	user := User{}
	body, err := io.ReadAll(request)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	return &user
}
