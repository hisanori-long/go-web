package main

import (
	"github.com/hisanori-long/go-web/my"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// main program.
func main() {
	my.Migrate()
}
