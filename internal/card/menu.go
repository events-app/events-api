package card

import (
	"errors"
	"fmt"
)

// Menu holds Name and wired card.
type Menu struct {
	Name string `json:"name"`
	Card string `json:"card"`
}

// menus stores all the menus
var menus = []Menu{
	Menu{Name: "Main menu", Card: "main"},
	Menu{Name: "Secured", Card: "secured"},
	Menu{Name: "Other menu", Card: "other"},
}

// GetAllMenus returns all menus
func GetAllMenus() (*[]Menu, error) {
	if len(menus) == 0 {
		return nil, fmt.Errorf("cannot find any object")
	}
	return &menus, nil
}

// FindMenu returns Content object based on name
func FindMenu(name string) (*Menu, error) {
	for _, c := range menus {
		if c.Name == name {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("cannot find any object of name: %s", name)
}

// AddMenu appends new Menu object
func AddMenu(name, card string) error {
	// Name of a menu must be unique for specyfic user.
	if !ValidateName(name) {
		return errors.New("name should be 4-30 characters long and should consists of letters, numbers, -, _")
	}
	if Find(name) != nil {
		return errors.New("menu with the name already exists")
	}
	menus = append(menus, Menu{Name: name, Card: card})

	return nil
}

// UpdateMenu changes Content object based on name
// Returns error if it was not found
func UpdateMenu(name, card string) error {
	for i := range menus {
		if menus[i].Name == name {
			menus[i].Card = card
			return nil
		}
	}
	return errors.New("menu not found")
}

func DeleteMenu(name string) error {
	if len(menus) == 0 {
		return errors.New("no menus in database")
	}
	var index int
	var found bool
	for i := range menus {
		if menus[i].Name == name {
			index = i
			found = true
		}
	}
	if !found {
		return errors.New("menu not found")
	}
	menus = append(menus[:index], menus[index+1:]...)
	return nil
}
