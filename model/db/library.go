package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Meta struct {
	Name       string  // title of library
	Type	     uint8   // currently can only be '0' (music)
}

type Directories struct {
	ID		     uint16  // unique numerical ID of dir
	Path	     string  // path to folder to scan
	Type       bool    // currently can only be '0' (absolute)
	Depth      uint8   // currently can only be NULL (traverse all)
}

type LibraryItems struct {
	// this is relatively empty for now
	// will contain more later
	ItemFile	 string  // directory-relative path to file
	Directory  uint16  // ID of directory from `directories` table
}