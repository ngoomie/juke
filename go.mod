module hyperboreal.zone/juke

go 1.20

// required by juke itself
require (
	github.com/BurntSushi/toml v1.3.2 // config
	github.com/gotk3/gotk3 v0.6.2 // GUI
	gorm.io/gorm v1.25.2-0.20230530020048-26663ab9bf55 // ORM
  gorm.io/driver/sqlite v1.5.2 // SQLite
)

// indirect requirements; used by juke's deps
require (
	github.com/jinzhu/inflection v1.0.0 // GORM
	github.com/jinzhu/now v1.1.5 // GORM
  github.com/mattn/go-sqlite3 v1.14.17 // SQLite
)