package database

import (
	"database/sql"

	"github.com/deividroger/ms-wallet/src/internal/entity"
)

type ClientDb struct {
	DB *sql.DB
}

func NewClientDb(db *sql.DB) *ClientDb {
	return &ClientDb{DB: db}
}

func (c *ClientDb) Get(id string) (*entity.Client, error) {

	client := entity.Client{}

	smt, err := c.DB.Prepare("SELECT id, name, email, created_at FROM clients WHERE id = ? ")

	if err != nil {

		return nil, err
	}

	defer smt.Close()

	row := smt.QueryRow(id)

	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {

		return nil, err
	}

	return &client, nil
}

func (c *ClientDb) Save(client *entity.Client) error {

	smt, err := c.DB.Prepare("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer smt.Close()

	_, err = smt.Exec(client.ID, client.Name, client.Email, client.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
