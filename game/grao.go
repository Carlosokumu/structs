package main

import "fmt"

//lets create some structs to play with
type attacker struct {
	attackpower int
	bonus       int
}
type sword struct {
	attacker
	twohanded bool
}

type gun struct {
	attacker
	bulletsremains int
}

//Lets create a method belonging to  two types gun and sword
func (s sword) Wield() bool {
	fmt.Printf("Wielding a sword")
	return true
}
func (g gun) Wield() bool {
	fmt.Printf("Wielding a gun")
	return true
}
func main() {

}
