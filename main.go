package main

import (
	"awesomeProject/bootstrap"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
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
	flag.Parse()
	path := strings.Join([]string{*userName, ":", *password, "@tcp(", *ip, ":", *port, ")/", *dbName, "?charset=utf8mb4"}, "")

	log.Println(path)
	if err := bootstrap.Bootstrap(path, *listen); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

}

func main() {
	fmt.Printf("begin service")
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	select {
	case <-ctx.Done():
	}
}
