package sqlite

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Meta struct {
	Name       string	// title of library
	Type	     uint8  // currently can only be '0' (music)
	Directory  string // full path to folder to scan
}

type AudioItems struct {
	// this is relatively empty for now
	// will contain more later
	File			 string // directory-relative file
}