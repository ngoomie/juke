module nu.hyperboreal.zone/juke

go 1.20

// required by juke itself
require (
	github.com/gotk3/gotk3 v0.6.2 // GUI
	github.com/mattn/go-sqlite3 v1.14.17 // SQLite
	github.com/pelletier/go-toml v1.9.5 // conf
	gopkg.in/reform.v1 v1.5.1 // ORM
)

require (
	github.com/AlekSi/pointer v1.1.0 // reform
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // reform
	github.com/pkg/errors v0.9.1 // reform
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c // reform
	golang.org/x/text v0.3.4 // reform
)
