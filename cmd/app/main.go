package main

import (
	"fmt"
	"log"
	"os"
	_"database/sql"
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
	fmt.Println("PATH:", db)
	//db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "ec2-52-86-56-90.compute-1.amazonaws.com",5432,"qbhewqmzoqeisd", "9a72432985934116cb32a55215dec1344005a2902da1539a1e2e261a30bf7486", "da2uvheuu1e2r1"))
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
	//log.Fatal(e.Start("localhost:8089"))
}
