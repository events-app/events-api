package main

// Content is datastructure for cards
type Content struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

var Cards = []Content{
	Content{Name: "main", Text: `# The New Event

	The New Event is the best event ever.
	You should definitelly attend!
	
	+ Register at [The New Event](http://thenewevent.com/).
	+ Come
	+ Have fun
	
	We are waiting for you!`},
	{Name: "secured", Text: "You are allowed to see this."},
}

func Find(name string) *Content {
	for _, c := range Cards {
		if c.Name == name {
			return &c
		}
	}
	return nil
}
