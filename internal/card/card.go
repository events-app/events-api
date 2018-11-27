package card

import "errors"

// Card holds unique key Name and content Text.
type Card struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

// cards stores all cards
var cards = []Card{
	Card{Name: "main", Text: `# The New Event

	The New Event is the best event ever.
	You should definitelly attend!
	
	+ Register at [The New Event](http://thenewevent.com/).
	+ Come
	+ Have fun
	
	We are waiting for you!`},
	Card{Name: "secured", Text: "You are allowed to see this."},
	Card{Name: "other", Text: "Other content"},
	Card{Name: "other2", Text: "Other content 2"},
}

// GetAll returns all cards
func GetAll() *[]Card {
	if len(cards) == 0 {
		return nil
	}
	return &cards
}

// Find returns Content object based on name
func Find(name string) *Card {
	for _, c := range cards {
		if c.Name == name {
			return &c
		}
	}
	return nil
}

// Add appends new Content object
func Add(name, text string) error {
	// Name of a card must be unique.
	if Find(name) != nil {
		return errors.New("card with the name already exists")
	}
	cards = append(cards, Card{Name: name, Text: text})

	return nil
}

// Update changes Content object based on name
// Returns error if it was not found
func Update(name, text string) error {
	for i := range cards {
		if cards[i].Name == name {
			cards[i].Text = text
			return nil
		}
	}
	return errors.New("card not found")
}
