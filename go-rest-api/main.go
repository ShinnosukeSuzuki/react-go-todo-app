package main

import (
	"go-rest-api/db"
	"go-rest-api/router"
)

func main() {
	db := db.NewDB()
	e := router.NewRouter(db)
	e.Logger.Fatal(e.Start(":8080"))
}
