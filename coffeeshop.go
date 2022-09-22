package main

import "fmt"

type CoffeeShop interface {
	Drink()
}

type Latte struct{}

type Americano struct{}

type Сappuccino struct{}

type ContextDrink struct {
	coffeeShop CoffeeShop
}

func (f *Latte) Drink() {
	fmt.Println("Latte")
}

func (f *Americano) Drink() {
	fmt.Println("Americano")
}

func (f *Сappuccino) Drink() {
	fmt.Println("Сappuccino")
}

func (c *ContextDrink) Drinkcoffee() {
	c.coffeeShop.Drink()
}

func NewContextEat(s CoffeeShop) *ContextDrink {
	return &ContextDrink{coffeeShop: s}
}
func (c *ContextDrink) SetDrink(s CoffeeShop) {
	c.coffeeShop = s
}

func main() {

	lat := NewContextEat(new(Latte))
	lat.Drinkcoffee()
	amr := NewContextEat(new(Americano))
	amr.Drinkcoffee()
	cap := NewContextEat(new(Сappuccino))
	cap.Drinkcoffee()
}
