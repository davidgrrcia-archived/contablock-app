package main

import (
	"github.com/davidevloper/HoldContract-go/cmd/graph"
	"github.com/davidevloper/HoldContract-go/cmd/server"
)

const port = ":3000"

func main() {

	//eth_gateway.CreateAbiBytecode()

	var appGraph graph.Graph = graph.NewGraph()
	appGraph.Build()

	server.Run(appGraph, port)
}
