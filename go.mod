module nu.hyperboreal.zone/juke

go 1.20

// required by juke itself
require (
	github.com/gotk3/gotk3 v0.6.2 // GUI
	github.com/pelletier/go-toml v1.9.5 // conf
	gorm.io/driver/sqlite v1.5.2 // SQLite
	gorm.io/gorm v1.25.2-0.20230530020048-26663ab9bf55 // ORM
)

// indirect requirements; used by juke's deps
require (
	// gorm.io/gorm
	github.com/jinzhu/inflection v1.0.0
	github.com/jinzhu/now v1.1.5
	// gorm.io/driver/sqlite
	github.com/mattn/go-sqlite3 v1.14.17
)