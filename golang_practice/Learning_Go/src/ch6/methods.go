package main

import "fmt"

type Dog struct {
	Breed string
	Weight int
	Sound string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d *Dog) SpeakMore() {
	d.Sound = fmt.Sprintf("%v! %v! %v!\n", d.Sound, d.Sound, d.Sound)
	fmt.Print(d.Sound)
}

func main() {

	poodle := Dog{"Poodle", 37, "Woof"}
	fmt.Println(poodle)
	poodle.Speak()
	
	poodle.Sound = "Arf"
	poodle.Speak()
	
	poodle.SpeakMore()
	poodle.SpeakMore()
	
}

/*
% go run methods.go
{Poodle 37 Woof}
Woof
Arf
Arf! Arf! Arf!
Arf! Arf! Arf!
! Arf! Arf! Arf!
! Arf! Arf! Arf!
!
*/