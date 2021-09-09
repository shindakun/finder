package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	findCol := flag.Int("find", 0, "CSV column to search from")
	rawCol := flag.Int("raw", 0, "CSV column to search in")

	flag.Parse()

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
		fMap[v[*findCol]] = true
	}

	for _, v := range o {
		oMap[v[*rawCol]] = true
	}

	for i := range fMap {
		if !oMap[i] {
			fmt.Println(i)
		}
	}
}
