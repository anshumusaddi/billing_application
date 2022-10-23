package main

import "github.com/anshumusaddi/billing_application/routes"

func main() {
	engine := routes.InitRoutes()
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
