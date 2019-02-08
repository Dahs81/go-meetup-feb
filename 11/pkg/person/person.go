package person

// Person -
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// NewPerson creates a new person
func (p Person) NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}
