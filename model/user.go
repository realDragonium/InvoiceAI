package model

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `pg:"pk,type:uuid,default:gen_random_uuid()"json:"username"`
	Name     string
	Email    string
	Password string
}
