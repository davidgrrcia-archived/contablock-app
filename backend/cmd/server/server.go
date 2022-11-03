package server

import (
	"fmt"
	"net/http"

	"github.com/davidevloper/HoldContract-go/cmd/graph"
)

func Run(graph graph.Graph, port string) {
	fmt.Println("Server running on port " + port)
	err := http.ListenAndServe(port, graph.GetRouter())
	if err != nil {
		panic(err)
	}
}
