package main

import (
	"fmt"
	"log"

	_ "database/sql"

	"github.com/bmstu-rsoi/rsoi-2022-lab1-ci-cd-AnastasiiaRumak/internal/person/repository"
	"github.com/bmstu-rsoi/rsoi-2022-lab1-ci-cd-AnastasiiaRumak/internal/person/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

const (
	dsn = "serverName=localhost;databaseName=test;user=testuser;password=testpassword"
)

func main() {

	/*
		listenAddr := os.Getenv("LISTEN_ADDR")
		if len(listenAddr) == 0 {
			listenAddr = ":8080"
		}
	*/
	dbConf := configure.NewLocal()
	//db, err := sqlx.Connect("postgres", dsn)
	db, err := sqlx.Connect("postgres", dbConf.GetDSN())
	if err != nil {
		log.Fatal(fmt.Errorf("error connecting to database: %w", err))
	}

	repo := repository.NewPG(db)
	uc := usecase.New(repo)
	handler := delivery.NewHandler(uc)

	e := echo.New()
	handler.Configure(e)

	//log.Fatsl(e.Start(address: "https://localhost:8890"))
	//log.Fatal(http.ListenAndServe(listenAddr, nil))
	log.Fatal(e.Start(configure.GetConnString()))
}
