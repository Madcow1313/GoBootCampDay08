package main

import (
	"fmt"
	"reflect"
	"strings"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

type AnotherStrangeBeing struct {
	Color    int
	IsAnimal bool
	Species  string
	Size     float64 `unit:"inches" funit:"m"`
}

func describePlant(p any) {
	values := reflect.ValueOf(p)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		tag := string(types.Field(i).Tag)
		tag = strings.ReplaceAll(tag, ":", "=")
		tag = strings.ReplaceAll(tag, "\"", "")
		if len(tag) != 0 {
			fmt.Printf("%v(%v):%v\n", types.Field(i).Name, tag, values.Field(i))
		} else {
			fmt.Printf("%v:%v\n", types.Field(i).Name, values.Field(i))
		}
	}
}

func main() {
	firstPlant := UnknownPlant{
		FlowerType: "rose",
		LeafType:   "tipycal",
		Color:      10655,
	}
	secondPlant := AnotherUnknownPlant{
		FlowerColor: 10,
		LeafType:    "lanceolate",
		Height:      15,
	}

	thirdSomething := AnotherStrangeBeing{
		Color:    5,
		IsAnimal: true,
		Species:  "probably penguin",
		Size:     1.32,
	}
	describePlant(firstPlant)
	fmt.Println()
	describePlant(secondPlant)
	fmt.Println()
	describePlant(thirdSomething)
}
