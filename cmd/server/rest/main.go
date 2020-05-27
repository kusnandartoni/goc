package main

import (
	"fmt"
	"github.com/kusnandartoni/goc/pkg/infrastructure/rest/router"
	"log"
	"net/http"
)

func main() {
	routersInit := router.InitRouter()
	endpoint := fmt.Sprintf(":%d", 8000) // TODO: need to move PORT to env config
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endpoint,
		Handler:        routersInit,
		ReadTimeout:    60, // TODO: need to move to env config
		WriteTimeout:   60, // TODO: need to move to env config
		MaxHeaderBytes: maxHeaderBytes,
	}

	fmt.Println("start http server listening " + endpoint)

	log.Fatal("%v", server.ListenAndServe())
}
