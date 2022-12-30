package main

import (
	"github.com/davidgrcia/tfg-informatica-nextjs/cmd/graph"
	"github.com/davidgrcia/tfg-informatica-nextjs/cmd/server"
	"github.com/joho/godotenv"
)

const port = ":4000"

func main() {

	//eth_gateway.CreateAbiBytecode()

	// load .env file
	err := godotenv.Load(".env.local")
	if err != nil {
		panic(err)
	}

	var appGraph = graph.NewGraph()
	appGraph.Build()

	server.Run(appGraph, port)
}
