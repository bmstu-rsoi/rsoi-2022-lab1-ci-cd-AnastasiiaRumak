package repository

import (
	"context"

	"github.com/bmstu-rsoi/rsoi-2022-lab1-ci-cd-AnastasiiaRumak/internal/models"
	"github.com/jmoiron/sqlx"
	"database/sql"
	"errors"
)
var (
	ErrNoPersonWithSuchID = errors.New("no person with such ID")
)

const (
	insertQuery = `INSERT INTO persons(name, age, address, work) VALUES($1, $2, $3, $4) RETURNING id`
	updateQuery     = `UPDATE persons SET name = $1, age = $2, address = $3, work = $4 WHERE id = $5`
	selectByIDQuery = `SELECT * FROM persons WHERE id = $1`
	selectAll = `SELECT * FROM persons`
	deleteQuery = `DELETE FROM persons WHERE id = $1`
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
	return id, nil
}


func (p *PG) DeletePerson(ctx context.Context, id int64) error {
	_, err := p.db.ExecContext(ctx, deleteQuery, id)
	return err
}

func (p *PG) GetPersonID(ctx context.Context, id int64) (models.Person, error) {
	row := p.db.QueryRowxContext(ctx, selectByIDQuery, id)

	var person BDlist
	err := row.StructScan(&person)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Person{}, ErrNoPersonWithSuchID
		}
		return models.Person{}, err
	}

	return toModel(person), nil
}

func (p *PG) GetAll(ctx context.Context) (*[]models.Person, error) {
	rows, err := p.db.QueryxContext(ctx, selectAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons = make([]models.Person, 0)
	for rows.Next() {
		var person BDlist
		err = rows.StructScan(&person)
		if err != nil {
			return nil, err
		}
		persons = append(persons, toModel(person))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &persons, nil
}

func (p *PG) UpdatePerson(ctx context.Context, person models.Person) error {
	_, err := p.db.ExecContext(ctx, updateQuery, person.Name, person.Age, person.Address, person.Work, person.ID)
	return err
}

type BDlist struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	Age     int64  `db:"age"`
	Address string `db:"address"`
	Work    string `db:"work"`
}

//func (b *BDlist) toModel() models.Person {
func toModel(b BDlist) models.Person {
	return models.Person{
		ID:      b.ID,
		Name:    b.Name,
		Age:     b.Age,
		Address: b.Address,
		Work:    b.Work,
	}
}