package contracts

import (
	"github.com/golang/concepts/interface/contracts/fruits"
	"github.com/golang/concepts/interface/contracts/vegetables"
)

type Salad interface {
	fruits.Fruit
	vegetables.Vegetable
}
