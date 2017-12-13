package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

var conv float64

type car struct {
	gasPedal      uint16 // vals 0-65535 (unassigned it 0 and up, where int16 goes to neg)
	brakePedal    uint16
	steeringWheel int16
	topSpeedKMH   float64
}

// There are two types of methods in GO (value receiver and pointer receivers)
// This is a method since its using a struct (i.e. class object)
// This is value reciver
// Makes a copy of the value you pass to it and thus, can mess around without anything touching it, where as pointer recivers will actually change the value
// Value rev
// func (c car) kmh() civer more efficent for smaller structs (becuase it has to make a copy), where for larger stucts, pointer revier is easier
float64 {
	conv = float64(c.gasPedal) * (c.topSpeedKMH / usixteenbitmax)
	return conv
}

func (c car) mph() float64 {
	// float64(c.gasPedal) is a type conversion
	return float64(c.gasPedal) * (c.topSpeedKMH / usixteenbitmax / kmhMultiple)
}

// Pointer receivers
func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKMH = newSpeed
}

func main() {
	aCar := car{gasPedal: 65000,
		brakePedal:    0,
		steeringWheel: 12561,
		topSpeedKMH:   225.0}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	aCar.newTopSpeed(500)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
}
