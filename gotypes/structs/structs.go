package main 

import (
    "github.com/lucjohnson/info344-class-code/gotypes/structs/person"
)

func main() {
    // prs := person.Person{FirstName: "Luc", LastName: "Johnson"}
    prs := person.NewPerson("Luc", "Johnson")
    prs.FirstName = "Mr."
    // fmt.Printf("%+v\n", person);
    
    // person.SayHello(&prs)
    // person.SayHello(prs)
    prs.SayHello()
}