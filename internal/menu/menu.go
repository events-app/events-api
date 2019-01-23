package menu

import (
	"fmt"
	"math/rand"

	"github.com/events-app/events-api/internal/card"
)

// Menu holds Name and wired card.
type Menu struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	CardId int    `json:"cardID"`
}

// menus stores all the menus
var menus = []Menu{
	Menu{ID: 1, Name: "Main menu", CardId: 1},
	Menu{ID: 2, Name: "Secured", CardId: 2},
	Menu{ID: 3, Name: "Other menu", CardId: 3},
}

// GetAll returns all menus
func GetAll() (*[]Menu, error) {
	if len(menus) == 0 {
		return nil, fmt.Errorf("cannot find any menu object")
	}
	return &menus, nil
}

// Get menu
func Get(id int) (*Menu, error) {
	for _, m := range menus {
		if m.ID == id {
			return &m, nil
		}
	}
	return nil, fmt.Errorf("cannot find a menu with ID: %d", id)
}

// Find returns menu object based on name
func Find(name string) (*Menu, error) {
	for _, m := range menus {
		if m.Name == name {
			return &m, nil
		}
	}
	return nil, fmt.Errorf("cannot find a menu named: %s", name)
}

// Add appends new Menu object
func Add(name string, cardId int) (int, error) {
	// Name of a menu must be unique for specyfic user.
	if !ValidateName(name) {
		return 0, fmt.Errorf("name should be 4-30 characters long and should consists of letters, numbers, -, _")
	}
	if m, _ := Find(name); m != nil {
		return 0, fmt.Errorf("menu with the name already exists")
	}
	id := rand.Intn(10000)
	menus = append(menus, Menu{ID: id, Name: name, CardId: cardId})

	return id, nil
}

// Update changes Menu object based on name
// Returns error if it was not found
func Update(id int, name string, cardId int) error {
	for i := range menus {
		if menus[i].ID == id {
			menus[i].Name = name
			menus[i].CardId = cardId
			return nil
		}
	}
	return fmt.Errorf("menu with id:%d not found", id)
}

// Delete menu
func Delete(id int) error {
	if len(menus) == 0 {
		return fmt.Errorf("no menus in database")
	}
	var index int
	var found bool
	for i := range menus {
		if menus[i].ID == id {
			index = i
			found = true
		}
	}
	if !found {
		return fmt.Errorf("menu not found")
	}
	menus = append(menus[:index], menus[index+1:]...)
	return nil
}

// GetCardOfMenu returns card of menu
func GetCardOfMenu(id int) (*card.Card, error) {
	for _, m := range menus {
		if m.ID == id {
			c, err := card.Get(m.CardId)
			if err != nil {
				return nil, fmt.Errorf("cannot any card connected to the menu")
			}
			return c, nil
		}
	}
	return nil, fmt.Errorf("cannot find a menu with ID: %d", id)
}
