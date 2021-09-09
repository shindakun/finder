package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("finder")

	find, err := os.Open("find.csv")
	if err != nil {
		panic(err)
	}
	out, err := os.Open("out.csv")
	if err != nil {
		panic(err)
	}

	findr := csv.NewReader(find)
	outr := csv.NewReader(out)

	f, err := findr.ReadAll()
	if err != nil {
		panic(err)
	}
	o, err := outr.ReadAll()
	if err != nil {
		panic(err)
	}

	fMap := make(map[string]bool)
	oMap := make(map[string]bool)

	for _, v := range f {
		fMap[v[0]] = true
	}

	for _, v := range o {
		oMap[v[0]] = true
	}

	for i := range fMap {
		if !oMap[i] {
			fmt.Println(i)
		}
	}
}
