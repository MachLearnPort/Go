package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal      uint16 // vals 0-65535 (unassigned it 0 and up, where int16 goes to neg)
	brakePedal    uint16
	steeringWheel int16
	topSpeedKMH   float64
}

// There are two types of methods in GO (value receiver and pointer receivers)
// This is a method since its using a struct (i.e. class object)
// This is value reciver
func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKMH / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKMH / usixteenbitmax) / kmhMultiple
}

func main() {
	aCar := car{gasPedal: 65000,
		brakePedal:    0,
		steeringWheel: 12561,
		topSpeedKMH:   225.0}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
}
