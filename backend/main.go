package main

import (
	"github.com/davidgrcia/tfg-informatica-nextjs/cmd/graph"
	"github.com/davidgrcia/tfg-informatica-nextjs/cmd/server"
)

const port = ":4000"

func main() {

	//eth_gateway.CreateAbiBytecode()

	var appGraph = graph.NewGraph()
	appGraph.Build()

	server.Run(appGraph, port)
}
