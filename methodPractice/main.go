package main

func main() {

	raj := Student{"Ankit", 7, "9D", "maths"}

}

type Student struct {
	Name    string
	Roll    int
	Class   string
	Subject string
}

func (v Student) GetClass() {

}
