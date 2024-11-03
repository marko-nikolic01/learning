package main

import (
	"fmt"
	"strconv"
)

type Person interface {
	toString() string
}

type Student struct {
	index     string
	firstName string
	lastName  string
	gpa       float64
}

type FootballPlayer struct {
	firstName string
	lastName  string
	position  string
}

func (student Student) toString() string {
	return "INDEX: " + student.index + ", NAME: " + student.firstName + " " + student.lastName + ", GPA: " + strconv.FormatFloat(student.gpa, 'f', 2, 64)
}

func (player FootballPlayer) toString() string {
	return "NAME: " + player.firstName + " " + player.lastName + ", POSITION: " + player.position
}

func main() {
	student := Student{"RA 69/2020", "Marko", "Nikolic", 9.03}
	player := FootballPlayer{"Cristiano", "Ronaldo", "LW"}

	fmt.Println(student.toString())
	fmt.Println(player.toString())
}
