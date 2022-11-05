package main

import (
	"github.com/davidgrcia/tfg-informatica-nextjs/cmd/graph"
	"github.com/davidgrcia/tfg-informatica-nextjs/cmd/server"
)

const port = ":3000"

func main() {

	//eth_gateway.CreateAbiBytecode()

	var appGraph graph.Graph = graph.NewGraph()
	appGraph.Build()

	server.Run(appGraph, port)
}
