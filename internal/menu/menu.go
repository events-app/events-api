package menu

import (
	"fmt"
	"time"

	"github.com/events-app/events-api/internal/card"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Menu holds Name and wired card.
type Menu struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CardID      int       `db:"card_id" json:"cardID"`
	DateCreated time.Time `db:"created" json:"dateCreated"`
	DateUpdated time.Time `db:"updated" json:"dateUpdated"`
}

// GetAll returns all menus
func GetAll(db *sqlx.DB) ([]Menu, error) {
	var menus []Menu
	const q = `SELECT * FROM menu`
	if err := db.Select(&menus, q); err != nil {
		return nil, errors.Wrap(err, "selecting menus")
	}
	return menus, nil
}

// Get menu
func Get(db *sqlx.DB, id int) (*Menu, error) {
	var menu Menu
	const q = `SELECT * FROM menu WHERE id = $1`
	if err := db.Get(&menu, q, id); err != nil {
		return nil, errors.Wrap(err, "selecting single menu")
	}
	return &menu, nil
}

// Find returns menu object based on name
func Find(db *sqlx.DB, name string) (*Menu, error) {
	var menu Menu
	const q = `SELECT * FROM menu WHERE name = $1`
	if err := db.Get(&menu, q, name); err != nil {
		return nil, errors.Wrap(err, "finding single menu")
	}
	return &menu, nil
}

// Add appends new Menu object
func Add(db *sqlx.DB, name string, cardID int, now time.Time) (*Menu, error) {
	// Name of a menu must be unique for specyfic user.
	if !ValidateName(name) {
		return nil, fmt.Errorf("name should be 4-30 characters long and should consists of letters, numbers, -, _")
	}
	// menu name must be unique
	if m, _ := Find(db, name); m != nil {
		return nil, fmt.Errorf("menu with the name already exists")
	}
	menu := Menu{
		Name:        name,
		CardID:      cardID,
		DateCreated: now.UTC(),
		DateUpdated: now.UTC(),
	}
	const q = `
		INSERT INTO menu
		(name, card_id, created, updated)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	rows, err := db.Query(q, menu.Name, menu.CardID, menu.DateCreated, menu.DateUpdated)
	if err != nil {
		return nil, errors.Wrap(err, "inserting menu")
	}
	var id int
	if rows.Next() {
		rows.Scan(&id)
	}
	menu.ID = id

	return &menu, nil
}

// Update changes Menu object based on name
// Returns error if it was not found
func Update(db *sqlx.DB, id int, name string, cardID int, now time.Time) error {
	menu, err := Get(db, id)
	if err != nil {
		return err
	}
	menu.Name = name
	menu.CardID = cardID
	menu.DateUpdated = now
	const q = `
		UPDATE menu SET
			name = $2,
			card_id = $3,
			updated = $4
		WHERE
			id = $1`
	_, err = db.Exec(q, id, menu.Name, menu.CardID, menu.DateUpdated)

	if err != nil {
		return errors.Wrap(err, "updating menu")
	}

	return nil
}

// Delete menu
func Delete(db *sqlx.DB, id int) error {
	const q = `DELETE FROM menu WHERE id = $1`
	if _, err := db.Exec(q, id); err != nil {
		return errors.Wrapf(err, "deleting menu %d", id)
	}

	return nil
}

// GetCardOfMenu returns card which is assigned to the specific menu
func GetCardOfMenu(db *sqlx.DB, id int) (*card.Card, error) {
	var card card.Card
	const q = `SELECT * 
				FROM card 
				WHERE id = (SELECT card_id FROM menu WHERE id = $1)`
	if err := db.Get(&card, q, id); err != nil {
		return nil, errors.Wrap(err, "selecting card")
	}
	return &card, nil
}
