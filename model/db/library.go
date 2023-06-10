package db

import (
	"gopkg.in/reform"
)

type LibraryItems struct {
	ItemFile	string	`reform:"item_file,pk`
	Artist		string	`reform:"artist"`
	Title			string	`reform:"title"`
	Album			string	`reform:"album"`
}