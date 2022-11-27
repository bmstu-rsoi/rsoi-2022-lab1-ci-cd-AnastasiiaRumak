package main

import (
	"fmt"
	"log"
	"os"
	_ "database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq" // ...

	"github.com/bmstu-rsoi/rsoi-2022-lab1-ci-cd-AnastasiiaRumak/internal/person/delivery"
	"github.com/bmstu-rsoi/rsoi-2022-lab1-ci-cd-AnastasiiaRumak/internal/person/repository"
	"github.com/bmstu-rsoi/rsoi-2022-lab1-ci-cd-AnastasiiaRumak/internal/person/usecase"
)

  

const (
	//dsn = "serverName=localhost;databaseName=test;user=postgres;password=postgres"
	//dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "postgres", "postgres", "postgres", 5432, "postgres")
)






func main() {
	//db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost",5432,"postgres", "postgres", "postgres"))
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(fmt.Errorf("error connecting to database: %w", err))
	}

	repo := repository.NewPG(db)
	uc := usecase.New(repo)
	handler := delivery.NewHandler(uc)
	
	e := echo.New()
	
	handler.Configure(e)
	//log.Fatal(e.Start("localhost:8089"))
	log.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}