package model

import "github.com/google/uuid"

type NewUser struct {
	tableName struct{} `pg:"users"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
}

type User struct {
	tableName struct{}  `pg:"users"`
	Id        uuid.UUID `pg:"pk_id,type:uuid"`
	Username  string
	Email     string
}

type DBUser struct {
	tableName struct{}  `pg:"users"`
	Id        uuid.UUID `pg:"pk_id,type:uuid, default:gen_random_uuid()"`
	Username  string
	Email     string
	Password  string
}
