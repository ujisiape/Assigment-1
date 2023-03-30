package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Name    string
	Address string
	Job     string
	Reason  string
}

var friends []Friend = []Friend{
	{
		Name:    "Rehon",
		Address: "Jalan Raya Kemang Village",
		Job:     "Mahasiswa",
		Reason:  "Karena saya suka ngoding",
	},
	{
		Name:    "Nugi",
		Address: "Jalan Raya Hang Tuah",
		Job:     "Mahasiswa",
		Reason:  "Karena saya suka Golang, cinta golang",
	},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run biodata.go [friend name]")
		return
	}

	name := os.Args[1]
	friend := findFriend(name)

	if friend == nil {
		fmt.Printf("%s is not found\n", name)
		return
	}

	fmt.Printf("Name: %s\nAddress: %s\nJob: %s\nReason: %s\n", friend.Name, friend.Address, friend.Job, friend.Reason)
}

func findFriend(name string) *Friend {
	for _, friend := range friends {
		if friend.Name == name {
			return &friend
		}
	}
	return nil
}
