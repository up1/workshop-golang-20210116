package cli_test

import (
	"cli"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockFileStore struct {
	filename string
	content  []byte
}

func (fs *mockFileStore) Read() []byte {
	return fs.content
}

func (fs *mockFileStore) Write(content []byte) {
	fs.content = content
}

func TestWriteDataSuccess(t *testing.T) {
	// Arrange
	store := mockFileStore{filename: "for testing"}
	userService := cli.UserService{Store: &store}

	// Act
	userService.AddNew("test name", 100)

	// Assert
	expected := `[{"id":0, "name":"test name", "age":100}]`
	res := userService.ListAll()
	assert.JSONEq(t, expected, res)
}
