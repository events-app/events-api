package card

import (
	"fmt"
	"math/rand"
)

// Card holds unique key Name and content Text.
type Card struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

// cards stores all cards
var cards = []Card{
	Card{ID: 1, Name: "main", Text: `# The New Event

	The New Event is the best event ever.
	You should definitelly attend!
	
	+ Register at [The New Event](http://thenewevent.com/).
	+ Come
	+ Have fun
	
	We are waiting for you!`},
	Card{ID: 2, Name: "secured", Text: "You are allowed to see this."},
	Card{ID: 3, Name: "other", Text: "Other content"},
	Card{ID: 4, Name: "other2", Text: "Other content 2"},
}

// GetAll returns all cards
func GetAll() (*[]Card, error) {
	if len(cards) == 0 {
		return nil, fmt.Errorf("cannot find any card object")
	}
	return &cards, nil
}

// Get one card
func Get(id int) (*Card, error) {
	for _, c := range cards {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("cannot find a card with ID: %d", id)
}

// Find returns Content object based on name
func Find(name string) (*Card, error) {
	for _, c := range cards {
		if c.Name == name {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("cannot find a card named: %s", name)
}

// Add appends new Content object
func Add(name, text string) (int, error) {
	// Name of a card must be unique for specyfic user.
	if !ValidateName(name) {
		return 0, fmt.Errorf("name should be 4-30 characters long and should consists of letters, numbers, -, _")
	}
	// card name must be unique
	if c, _ := Find(name); c != nil {
		return 0, fmt.Errorf("card with the name already exists")
	}
	id := rand.Intn(10000)
	cards = append(cards, Card{ID: id, Name: name, Text: text})

	return id, nil
}

// Update changes Card object based on name
// Returns error if it was not found
func Update(id int, name, text string) error {
	for i := range cards {
		if cards[i].ID == id {
			cards[i].Name = name
			cards[i].Text = text
			return nil
		}
	}
	return fmt.Errorf("card with id:%d not found", id)
}

// Delete card
func Delete(id int) error {
	if len(cards) == 0 {
		return fmt.Errorf("no cards in database")
	}
	var index int
	var found bool
	for i := range cards {
		if cards[i].ID == id {
			index = i
			found = true
		}
	}
	if !found {
		return fmt.Errorf("card not found")
	}
	cards = append(cards[:index], cards[index+1:]...)
	return nil
}
