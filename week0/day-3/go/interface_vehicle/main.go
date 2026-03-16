package main

import "fmt"

type controller interface {
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed       int
	enginePower int
}

type bicycle struct {
	speed      int
	humanPower int
}

func (v *vehicle) speedUp() int {
	// TODO
	return 0
}
func (v *vehicle) speedDown() int {
	// TODO
	return 0
}
func (b *bicycle) speedUp() int {
	// TODO
	return 0
}
func (b *bicycle) speedDown() int {
	// TODO
	return 0
}

func speedUpAndDown(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}

func main() {
	v := &vehicle{enginePower: 5}
	b := &bicycle{humanPower: 5}
	speedUpAndDown(v)
	speedUpAndDown(b)
}
