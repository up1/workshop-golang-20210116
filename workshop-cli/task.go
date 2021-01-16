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
	// Read file
	allUsers := us.Store.Read()
	if len(allUsers) == 0 {
		allUsers = []byte("[]")
	}
	var out Users
	err := json.Unmarshal(allUsers, &out)
	if err != nil {
		fmt.Println("Error> ", err)
		return
	}
	out = append(out, u1)
	content, _ := json.Marshal(out)
	us.Store.Write(content)
}

// ListAll : List all users from file
func (us UserService) ListAll() string {
	allUsers := us.Store.Read()
	var out Users
	err := json.Unmarshal(allUsers, &out)
	if err != nil {
		fmt.Println("Error ", err)
	}
	b, err := json.MarshalIndent(out, "", "    ")
	if err != nil {
		fmt.Println("Error ", err)
		return ""
	}
	return string(b)
}
