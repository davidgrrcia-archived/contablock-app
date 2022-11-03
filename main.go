package main

import (
	"github.com/davidevloper/HoldContract-go/backend/cmd/graph"
	"github.com/davidevloper/HoldContract-go/backend/cmd/server"
	eth_gateway "github.com/davidevloper/HoldContract-go/eth-gateway"
)

const port = ":8080"

func main() {

	eth_gateway.CreateAbiBytecode()

	var appGraph graph.Graph = graph.NewGraph()
	appGraph.Build()

	server.Run(appGraph, port)
}
