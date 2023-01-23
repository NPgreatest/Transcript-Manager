package bootstrap

import (
	"awesomeProject/lib/db"
	"awesomeProject/route"
	"log"
	"net/http"
)

func Bootstrap(path, addr string) error {
	if err := db.ConnectDB(path); err != nil {
		return err
	}
	route.RegisterRoutes()
	log.Println("listen " + addr)
	return http.ListenAndServe(addr, nil)
}
