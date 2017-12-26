package models

import (
	"time"
)
/*
	--------------------------------------
	NOTE : bila ingin menambahkan model
	tambahkan masing masing struct dengan
	nama model & modelResult ..
 	--------------------------------------
*/
type (
	Pet struct {
		Id uint `json:"id"`
		Name string `form:"name" json:"name" binding:"required"`
		Age int `form:"age" json:"age" binding:"required"`
		Photo string `form:"photo" json:"photo" binding:"required"`
	}

	PetResult struct {
		Id uint `json:"id"`
		Name string `json:"name"`
		Age int `json:"age"`
		Photo string `json:"photo"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
