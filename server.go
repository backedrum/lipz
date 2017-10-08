package main

import (
	"github.com/backedrum/lipz/api"
	"github.com/backedrum/lipz/models"
	"github.com/backedrum/lipz/routes"
	"github.com/urfave/negroni"
)

func main() {
	db := models.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":3000")
}
