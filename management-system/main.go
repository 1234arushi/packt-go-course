package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/packt-go-course/management-system/contact"
)

func displayHelp() {
	fmt.Println("Available commands:")
	fmt.Println("add <name> <phone_number> <email> - Add a new contact")
	fmt.Println("view <name> - View contact details")
	fmt.Println("delete <name> - Delete a contact")
	fmt.Println("list - List all contacts")
	fmt.Println("help - Display this help message")
	fmt.Println("quit - Exit the program")

}

func main() {
	fmt.Println("Welcome to the Contacts Management System!")
	displayHelp()
	//creates a reader that listens to what you type in the terminal
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n Enter a command: ")
		scanner.Scan()            //waits for the user to type a line and press enter
		command := scanner.Text() // i/p converted to text
		//splts the line into words using spaces
		args := strings.Fields(command) // fields splits the string around each instance of one or more consecutive white space characters
		if len(args) == 0 {
			fmt.Println("Please enter a valid command.")
		}
		switch args[0] {
		case "add":
			if len(args) < 4 {
				fmt.Println("Usage : add <name> <phone_number> <email>")
				continue
			}
			name := args[1]
			phoneNo := args[2]
			email := args[3]
			contact := contact.Contact{
				Name:    name,
				PhoneNo: phoneNo,
				Email:   email,
			}
			err := contact.AddContact(contact)
			if err != nil {
				fmt.Println("Error : ", err)
			} else {
				fmt.Println("Contact added successfully!")
			}
		case "view":
			if len(args) < 2 {
				fmt.Println("Usage : view <name>")
				continue
			}
			name := args[1]
			contact, err := contact.ViewContact(name)
			if err != nil {
				fmt.Println("Error : ", err)
			} else {
				fmt.Println("Contact Details : ", contact)
			}
		case "delete":
			if len(args) < 2 {
				fmt.Println("Usage : delete <name>")
				continue
			}
			name := args[1]
			err := contact.DeleteContact(name)
			if err != nil {
				fmt.Println("Error : ", err)
			} else {
				fmt.Println("Contact deleted successfully!")
			}
		case "help":
			displayHelp()
		case "list":
			contacts := contact.GetAllContacts()
			if len(contacts) == 0 {
				fmt.Println("No contacts found.")
			} else {
				fmt.Println("List of contacts : ")
				for _, c := range contacts {
					fmt.Printf("Name : %s, Phone : %s, Email : %s\n", c.Name, c.PhoneNo, c.Email)
				}
			}
		case "quit":
			fmt.Println("Exiting the program. Goodbye!")
			os.Exit(0) //stops the program and 0 -> closed successfully
		default:
			fmt.Println("Unknown command. Type 'help' for a list of available commands.")

		}
	}

}
