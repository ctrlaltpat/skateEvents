package models

import (
	"database/sql"
)

type SkateModel struct {
	DB *sql.DB
}

type Skate struct {
	Id       int    `json:"id"`
	SkaterId int    `json:"skater_id" binding:"required"`
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Brand    string `json:"brand" binding:"required,min=3,max=100"`
	Plates   string `json:"plates" binding:"required,min=3,max=100"`
	Wheels   string `json:"wheels" binding:"required,min=3,max=100"`
}
