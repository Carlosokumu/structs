package main

import "fmt"

//struct-goes(syntax) by 'type(keyword)->name(name(struct))->struct(keyword)'
type Person struct {
	firstname  string
	secondname string
}

/*
    METHOD
   context=go -> goes by: A method is a function with a special receiver argument.
   In this example, the getName() method has a receiver of type Person named p.
*/
func (p Person) getName() string {
	return p.firstname
}

/*
   Pointer receivers
   This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
*/
func (p *Person) modify(firstname string, secondname string) {
	p.firstname = firstname
	p.secondname = secondname
}
func (p Person) modify2(firstname string, secondname string) {
	p.firstname = firstname
	p.secondname = secondname
}

func main() {
	//instance
	v := Person{"carlos", "okumu"}
	fmt.Println(v.firstname)
	//pointer
	p := &v
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(v.getName())
	v.modify("Reizger", "ombaso")
	//name altered->confirm
	fmt.Println(v.getName())
	//another instance-confirm modification inability of modify2()
	m := Person{"Adams", "nyabs"}
	m.modify2("Asley", "oogs")
	fmt.Println(m.getName()) //output->"Adams"

}
