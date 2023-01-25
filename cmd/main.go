package main

import (
	"ShortLink/internal/config"
	"ShortLink/internal/service"
	"flag"
)

func main() {
	dbStorageMode := flag.Bool("use_db_storage", false, "use for save data in PostgreSQL storage")
	flag.Parse()
	cfg := config.GetConfig(*dbStorageMode)
	srv := service.NewService(cfg)

	srv.Logger.Fatal(srv.Start())

}
