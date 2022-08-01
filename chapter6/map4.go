package chapter6

import "fmt"

func Chapter6main8() {
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

	// name, ok := elements["Li"]
	// fmt.Println(name, ok)

	if name, ok := elements["Li"]; ok {
		fmt.Println(name, ok)
	} //循环

	// for keys which don't exist, the zero type is returned
	// for strings it is the empty string
	// fmt.Println(elements["Un"])
}
