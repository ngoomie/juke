package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

//! don't use this yet

type Items struct {
	ItemType	uint8 // currently can only be '0' (ind. audio file)
}