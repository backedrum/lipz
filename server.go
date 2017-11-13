package main

import (
	"github.com/backedrum/lipz/routes"
	"github.com/urfave/negroni"
)

func main() {
	routes := routes.NewRoutes()
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":3000")
}
