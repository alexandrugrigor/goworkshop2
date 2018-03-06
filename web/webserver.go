package web

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"fmt"
	"time"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Handle(route.Pattern, log(route.HandlerFunc)).Methods(route.Method)
	}
	var port = getPort()
	fmt.Println("Starting server on port:" + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(err)
	}
}

func log(routeFunc RouteFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request to: " + r.URL.Path)
		start := time.Now().UnixNano()
		err := routeFunc(w, r) // call original
		//handle the errors
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Errorf("error occurred while processing the request:%v", err)
		}
		end := time.Now().UnixNano()
		fmt.Printf("Request took: %d nano\n", (end - start))
	})
}

func getPort() string {
	var port = os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}
