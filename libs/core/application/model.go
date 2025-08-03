package application

import "github.com/uptrace/bun"

type Application struct {
	bun.BaseModel `bun:"table:applications"`
	ID            string `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Name          string `bun:",type:text,notnull,unique:name_and_version" json:"name"`
	Version       string `bun:",type:text,notnull,unique:name_and_version" json:"version"`
	Description   string `bun:",type:text" json:"description"`
}
