package cli

import (
	"encoding/json"
	"fmt"
)

type Users []User

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserService struct {
	Store Store
}

// AddNew : Add new user to file
func (us UserService) AddNew(name string, age int) {
	u1 := User{Name: name, Age: age}
	// Read data from file
	allUsers := us.Store.Read()
	if len(allUsers) == 0 {
		allUsers = []byte("[]")
	}
	// JSON message to Struct
	out := jsontoStruct(allUsers)
	// Add new user
	out = append(out, u1)
	// Struct to JSON message
	content, _ := json.Marshal(out)
	// Save data to file
	us.Store.Write(content)
}

// ListAll : List all users from file
func (us UserService) ListAll() string {
	// Read data from file
	allUsers := us.Store.Read()
	// JSON message to Struct
	users := jsontoStruct(allUsers)
	// Formatting output
	return formatting(users)
}

func jsontoStruct(content []byte) Users {
	var out Users
	err := json.Unmarshal(content, &out)
	if err != nil {
		fmt.Println("Error> ", err)
		return nil
	}
	return out
}

func formatting(users Users) string {
	b, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		fmt.Println("Error ", err)
		return ""
	}
	return string(b)
}
