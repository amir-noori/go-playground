package simpleDataStructure

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type User struct {
	name string
	age  int
}

type Account struct {
	id   int
	user User
}

func Run() {
	user := createUser("Amir", 34)
	fmt.Printf("User name: %s\n", (*user).name)

	account := createAccount(user)
	fmt.Printf("Account id: %d\n", (*account).id)

	fileContent := readFile("cars.json")
	fmt.Println("file content: " + *fileContent)
}

func createUser(name string, age int) *User {
	user := User{name: name, age: age}
	fmt.Printf("User created with name: %s\n", user.name)
	return &user
}

func createAccount(user *User) *Account {
	account := Account{id: 100, user: *user}
	fmt.Printf("Account created with user name: %s\n", (*user).name)
	return &account
}

func readFile(fileName string) *string {
	result := ""

	fmt.Println("reading file: " + fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("cannot open file" + err.Error())
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("error reading line: " + err.Error())
		} else {
			result = result + line
		}
	}

	return &result
}
