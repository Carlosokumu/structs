package main

import "fmt"

func main() {
	//Make a map that will use String type as the key then Int as the value
	x := make(map[string]int)
	x["KEY"] = 10
	fmt.Println(x["KEY"])
	//Shorthand maps.
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	ok, name := elements["Un"]
	fmt.Println(name, ok)
	fmt.Println(ok)
	fmt.Println(name)

}
