package bootstrap

import (
	"awesomeProject/lib/db"
	"log"
	"net/http"
)

func Bootstrap(path, addr string) error {
	if err := db.ConnectDB(path); err != nil {
		return err
	}
	log.Println("listen " + addr)
	return http.ListenAndServe(addr, nil)
}
