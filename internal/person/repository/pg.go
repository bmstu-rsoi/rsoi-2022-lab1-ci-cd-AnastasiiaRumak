package repository

import (
	"context"

	"github.com/bmstu-rsoi/rsoi-2022-lab1-ci-cd-AnastasiiaRumak/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	insertQuery = `INSERT INTO persons(name, age, address, work) VALUES($1, $2, $3, $4) RETURNING id`
)

type PG struct {
	db *sqlx.DB
}

func NewPG(db *sqlx.DB) *PG {
	return &PG{db: db}
}

func (p *PG) CreatePerson(ctx context.Context, person models.Person) (int64, error) {
	row := p.db.QueryRowContext(ctx, insertQuery, person.Name, person.Age, person.Address, person.Work)
	var id int64
	err := row.Scan(&id)
	row.Scan(&id)

	if err != nil {
		return 0, err
	}
	/*if err := row.Err();
	err != nil {
		return id, nil
	}*/

	return id, nil

}
