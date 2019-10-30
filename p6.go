package main

import (
	"fmt"
)

const uxixteenbitmax float64 = 65535
const kmh_miles float64 = 1.60934

type car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	steering_wheel uint16
	top_speed      float64
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed / uxixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed / uxixteenbitmax / kmh_miles)
}

func main() {
	a_car := car{gas_pedal: 22112,
		brake_pedal:    0,
		steering_wheel: 12231,
		top_speed:      255.0}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

}
