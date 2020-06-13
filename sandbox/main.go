package main

import "fmt"

type Bet struct {
	wagerAmount uint
}

func findUser(db string) func(string) string {
	return func(userID string) string {
		user := fmt.Sprintf("%s %s", db, userID)
		return user
	}
}

func WithStatus(db string) func(status string) string {
	return func(status string) string {
		str := fmt.Sprintf("%s %s", db, status)
		return str
	}
}

func main() {
	find := findUser("database")
	str := find("1234")

	fmt.Println(str)
}
