package main

import (
	"awesomeProject/lib/db"
	"awesomeProject/route"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	listen   = flag.String("listen", "localhost:8080", "listen")
	userName = flag.String("userName", "root", "db userName")
	password = flag.String("password", "root", "db password")
	ip       = flag.String("ip", "localhost", "db host")
	port     = flag.String("port", "3306", "db port")
	dbName   = flag.String("dbName", "gpa_manage", "db Name")
)

func init() {
	//flag.Parse()
	//path := strings.Join([]string{*userName, ":", *password, "@tcp(", *ip, ":", *port, ")/", *dbName, "?charset=utf8mb4"}, "")

	//log.Println(path)
	//if err := bootstrap.Bootstrap(path, *listen); err != nil {
	//	log.Println(err.Error())
	//	os.Exit(1)
	//}

}

func main() {
	fmt.Printf("begin service")
	engine := gin.New()
	route.Register(engine)
	addr := fmt.Sprintf("%s", *listen)

	srv := http.Server{
		Addr:    addr,
		Handler: engine,
	}
	path := strings.Join([]string{*userName, ":", *password, "@tcp(", *ip, ":", *port, ")/", *dbName, "?charset=utf8mb4"}, "")
	if err := db.ConnectDB(path); err != nil {
		fmt.Println("cannot connect MYSQL")
		return
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				fmt.Println("Server closed!")
				return
			} else {
				panic(err)
			}
		}
	}()

	dealSignal(&srv)
}
func dealSignal(server *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case sig := <-sigChan:
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("Server is going to shutdown...")
			c, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
			defer cancelFunc()

			if err := server.Shutdown(c); err != nil {
				fmt.Println("Stop server error:%v", err)
			}
			return
		}
	}

}
