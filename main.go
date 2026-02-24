package main

import "fmt"

type Vehicle struct {
	Brand string
	Model string
	Year  int
	Color string
}

type VehicleError struct {
	Err         error
	VehicleType string //as the additional information
}

func (ve VehicleError) Error() string {
	return fmt.Sprintf("%s vehicle error : %v", ve.VehicleType, ve.Err)
}

type VehicleInterface interface {
	Start()
	Stop()
	Steer()
}
type Car struct {
	Vehicle
	NumDoors   int
	EngineType string
}

type Boat struct {
	Vehicle
	Length         int
	PropulsionType string
}

type MotorCycle struct {
	Vehicle
	NumWheels  int
	HasSideCar bool
}

func (c Car) Start() {
	fmt.Printf("Starting the %s %s with %d doors\n", c.Color, c.Brand, c.NumDoors)
}
func (c Car) Stop() {
	fmt.Printf("Stopping the %s %s with %d doors\n", c.Color, c.Brand, c.NumDoors)
}
func (c Car) Steer() {
	fmt.Printf("Steering the %s %s with %d doors\n", c.Color, c.Brand, c.NumDoors)
}

// boat methods
func (b Boat) Start() {
	fmt.Printf("Starting the %s boat with %d length\n", b.Color, b.Length)
}
func (b Boat) Stop() {
	fmt.Printf("Stopping the %s boat with %d length\n", b.Color, b.Length)
}
func (b Boat) Steer() {
	fmt.Printf("Steering the %s boat with %d length\n", b.Color, b.Length)
}

// boat methods
func (m MotorCycle) Start() {
	fmt.Printf("Starting the %s motorcycle with %d wheels\n", m.Brand, m.NumWheels)
	if m.NumWheels < 0 {
		panic("negative number of wheels")
	}
}
func (m MotorCycle) Stop() {
	fmt.Printf("Stopping the %s motorcyle with %d wheels\n", m.Brand, m.NumWheels)
}
func (m MotorCycle) Steer() {
	fmt.Printf("Steering the %s motorcyles with %d wheels\n", m.Brand, m.NumWheels)
}

func main() {

}
