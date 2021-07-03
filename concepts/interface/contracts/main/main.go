package main

import (
	"github.com/golang/concepts/interface/contracts"
	"github.com/golang/concepts/interface/contracts/fruits"
	"github.com/golang/concepts/interface/contracts/vegetables"
)

func FruitSalad(fruit fruits.Fruit) {
	fruit.Printname()
}

func VegetableSalad(vege vegetables.Vegetable) {
	vege.Printname()
}

func GetSalad(ingrediants ...contracts.Salad) {
	for _, ingrediant := range ingrediants {
		ingrediant.Printname()
	}
}

func main() {
	FruitSalad(fruits.Orange{})
	FruitSalad(fruits.Apple{})

	VegetableSalad(vegetables.Potato{})
	VegetableSalad(vegetables.Tomato{})

	GetSalad(fruits.Orange{})
	GetSalad(fruits.Apple{})

	GetSalad(vegetables.Potato{})
	GetSalad(vegetables.Tomato{})

	GetSalad(vegetables.Tomato{}, vegetables.Potato{}, fruits.Orange{}, fruits.Apple{})
}
