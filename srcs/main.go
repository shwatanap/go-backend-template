package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/shwatanap/go-backend-template/srcs/infra/sql"
	"github.com/shwatanap/go-backend-template/srcs/presen/handler"
	"github.com/shwatanap/go-backend-template/srcs/wire"
)

func main() {
	db := sql.NewDriver()
	userHandler := wire.InitTemplateHandler(db)

	handler.InitRouting(userHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr: ":" + port,
	}
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln("Server closed with error:", err)
	}
}
