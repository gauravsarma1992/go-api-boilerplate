package server

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model

		Username       string `json:"username"`
		Password       string `json:"password" gorm:"-"`
		HashedPassword string `json:"-"`

		Name   string `json:"name"`
		Email  string `json:"email"`
		Mobile string `json:"mobile"`
	}
)
