package main

import (
	"github.com/darienkentanu/test-erajaya/config"
	"github.com/darienkentanu/test-erajaya/routes"
)

func main() {
	db := config.InitDB()
	e := routes.New(db)

	e.Logger.Fatal(e.Start(":8000"))
}
