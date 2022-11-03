package server

import (
	"net/http"

	"github.com/davidevloper/HoldContract-go/backend/cmd/graph"
)

func Run(graph graph.Graph, port string) {
	err := http.ListenAndServe(port, graph.GetRouter())
	if err != nil {
		panic(err)
	}
}
