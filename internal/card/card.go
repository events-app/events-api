package card

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Card holds unique key Name and content Text.
type Card struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Text        string    `json:"text"`
	DateCreated time.Time `db:"created" json:"dateCreated"`
	DateUpdated time.Time `db:"updated" json:"dateUpdated"`
}

// GetAll returns all cards
func GetAll(db *sqlx.DB) ([]Card, error) {
	var cards []Card
	const q = `SELECT * FROM card`
	if err := db.Select(&cards, q); err != nil {
		return nil, errors.Wrap(err, "selecting cards")
	}
	return cards, nil
}

// Get one card
func Get(db *sqlx.DB, id int) (*Card, error) {
	var card Card
	const q = `SELECT * FROM card WHERE id = $1`
	if err := db.Get(&card, q, id); err != nil {
		return nil, errors.Wrap(err, "selecting single card")
	}
	return &card, nil
}

// Find returns Content object based on name
func Find(db *sqlx.DB, name string) (*Card, error) {
	var card Card
	const q = `SELECT * FROM card WHERE name = $1`
	if err := db.Get(&card, q, name); err != nil {
		return nil, errors.Wrap(err, "finding single card")
	}
	return &card, nil
}

// Add appends new Content object
func Add(db *sqlx.DB, name, text string, now time.Time) (*Card, error) {
	// Name of a card must be unique for specyfic user.
	if !ValidateName(name) {
		return nil, fmt.Errorf("name should be 4-30 characters long and should consists of letters, numbers, -, _")
	}
	// card name must be unique
	if c, _ := Find(db, name); c != nil {
		return nil, fmt.Errorf("card with the name already exists")
	}
	card := Card{
		Name:        name,
		Text:        text,
		DateCreated: now.UTC(),
		DateUpdated: now.UTC(),
	}
	const q = `
		INSERT INTO card
		(name, text, created, updated)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	rows, err := db.Query(q, card.Name, card.Text, card.DateCreated, card.DateUpdated)
	if err != nil {
		return nil, errors.Wrap(err, "inserting card")
	}
	var id int
	if rows.Next() {
		rows.Scan(&id)
	}
	card.ID = id
	
	return &card, nil
}

// Update changes Card object based on name
// Returns error if it was not found
func Update(db *sqlx.DB, id int, name, text string, now time.Time) error {
	card, err := Get(db, id)
	if err != nil {
		return err
	}
	card.Name = name
	card.Text = text
	card.DateUpdated = now
	const q = `
		UPDATE card SET
			name = $2,
			text = $3,
			updated = $4
		WHERE
			id = $1`
	_, err = db.Exec(q, id, card.Name, card.Text, card.DateUpdated)

	if err != nil {
		return errors.Wrap(err, "updating card")
	}

	return nil
}

// Delete card
func Delete(db *sqlx.DB, id int) error {
	const q = `DELETE FROM card WHERE id = $1`
	if _, err := db.Exec(q, id); err != nil {
		return errors.Wrapf(err, "deleting card %d", id)
	}

	return nil
}
