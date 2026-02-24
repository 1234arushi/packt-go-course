package contact

import (
	"errors"

	"github.com/packt-go-course/management-system/validation"
)

type Contact struct {
	Name    string
	PhoneNo string
	Email   string
}

var contactsDB map[string]Contact //simulating a database with a map

func init() {
	contactsDB = make(map[string]Contact)

}
func (c Contact) AddContact(contact Contact) error {
	if contact.Name == "" {
		return errors.New("Contact name cannot be empty")
	}
	if _, exists := contactsDB[contact.Name]; exists {
		return errors.New("Contact already exsists")
	}
	if !validation.IsValidPhone(contact.PhoneNo) {
		return errors.New("Invalid phone number format")
	}
	if !validation.IsValidEmail(contact.Email) {
		return errors.New("Invalid email format")
	}
	contactsDB[contact.Name] = contact
	return nil

}

func ViewContact(name string) (Contact, error) {
	contact, exists := contactsDB[name]
	if !exists {
		return Contact{}, errors.New("contact not found")
	}
	return contact, nil
}

func DeleteContact(name string) error {
	_, exists := contactsDB[name]
	if !exists {
		return errors.New("Contact not found")
	}
	delete(contactsDB, name)
	return nil
}

func GetAllContacts() []Contact {
	contacts := make([]Contact, 0, len(contactsDB))
	for _, contact := range contactsDB {
		contacts = append(contacts, contact)
	}
	return contacts
}
