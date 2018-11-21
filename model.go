package main

// Content is datastructure for cards
type Content struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

var cards = []Content{
	Content{Name: "main", Text: `# The New Event

	The New Event is the best event ever.
	You should definitelly attend!
	
	+ Register at [The New Event](http://thenewevent.com/).
	+ Come
	+ Have fun
	
	We are waiting for you!`},
	Content{Name: "secured", Text: "You are allowed to see this."},
	Content{Name: "other", Text: "Other content"},
	Content{Name: "other2", Text: "Other content 2"},
}

// GetAll returns all cards
func GetAll() *[]Content {
	if len(cards) == 0 {
		return nil
	}
	return &cards
}

// Find returns Content object based on name
func Find(name string) *Content {
	for _, c := range cards {
		if c.Name == name {
			return &c
		}
	}
	return nil
}

// Add appends new Content object
func Add(name, text string) {
	cards = append(cards, Content{Name: name, Text: text})
}

// Update changes Content object based on name
// Returns true if item updated, and false if it was not found
func Update(name, text string) bool {
	for i := range cards {
		if cards[i].Name == name {
			cards[i].Text = text
			return true
		}
	}
	return false
}
