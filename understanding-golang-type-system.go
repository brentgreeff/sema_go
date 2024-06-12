package main

import "fmt"

func main() {
	//child := Child{Base: Base{ID: id}, a: a, b: b}

	p := Speaker {
		Person: Person {
			name: "Brent",
			age: 42,
			city: "Cologne",
			phone: "+49",
		},
		//speaksOn ['today', 'tomorrow']
	}
	var m = Meeting{}
	m.people := [...]Person{p}

	for _, v := range m.people {
		v.SayHello()
		v.GetDetails()
	}
	// fmt.Println("Person", p)
}

type Meeting struct {
	location string
	city string
	date time.Time
	people []People
}

type People interface {
	SayHello()
	GetDetails()
}

type People interface {
	SayHello()
	GetDetails()
}

type Person struct {
	name string
	age int
	city,phone string
}

type Speaker struct {
	Person
	speaksOn []string
	pastEvents []string
}

type Organizer struct {
	Person
	meetups []string
}

type Attendee struct {
	Person
	interests []string
}

func (p Person) sayHello() {
	fmt.Printf("Hi, I am %s from %s\n", p.name, p.city)
}

func (p *Person) incAge() {
	p.age++
}

func (p Person) getDetails() {
	fmt.Printf("I am %d", p.age)
}
