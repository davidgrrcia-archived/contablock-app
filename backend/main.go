package main

import (
	"os"

	"github.com/davidgrcia/tfg-informatica-nextjs/cmd/graph"
	"github.com/davidgrcia/tfg-informatica-nextjs/cmd/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic(err)
	}

	port := os.Getenv("BACKEND_PORT")

	var appGraph = graph.NewGraph()
	appGraph.Build()

	server.Run(appGraph, ":"+port)
}
