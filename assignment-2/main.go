package main

import (
	db "assignment-2/database"
	"assignment-2/routers"
)

func main() {
	db.Init()
	routers.ServerOn().Run(":8080")
}
