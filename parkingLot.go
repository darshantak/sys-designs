// Design a car parking application
// 1.There are two categories of cars: SUV and hatchback.
// 2.Maintain a count of how many SUV and hatchback cars enter the premise.
// 3.Calculate the payment each car has to make based upon the rates as hatchback parking as 10 rupees per hour and SUV being 20 rupees an hour.
// 4.In case if hatchback occupancy is full then hatchback cars can occupy SUV spaces but with hatchback rates.
// 5.During exit there needs to be the system to inform the user how much they have to pay.
// 6.Admin can see all the cars which are parked in the system.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CarType string

const (
	Suv       CarType = "suv"
	Hatchback CarType = "hatchback"
)

type Car struct {
	Name        string
	Type        CarType
	CheckInTime time.Time
}
type CarParking struct {
	suvCars map[string]Car
	hbCars  map[string]Car
	maxOccu int
}

func NewCarParking() *CarParking {
	return &CarParking{
		suvCars: make(map[string]Car),
		hbCars:  make(map[string]Car),
		maxOccu: 10,
	}
}

func (cp *CarParking) parkVehicle(name string, carType CarType) {
	if carType == "suv" {
		for _, v := range cp.suvCars {
			if v.Name == name {
				fmt.Println("Car already exists")
				return
			}
		}
		if len(cp.suvCars) == cp.maxOccu {
			fmt.Println("Max occupancy reached")
			return
		} else {
			newCar := Car{Name: name, Type: CarType(carType), CheckInTime: time.Now()}
			cp.suvCars[name] = newCar
			fmt.Println("SUV car parked")
			return
		}
	} else {
		for _, v := range cp.hbCars {
			if v.Name == name {
				fmt.Println("Car already exists")
				return
			}
		}
		if len(cp.hbCars) == cp.maxOccu {
			if len(cp.suvCars) < cp.maxOccu {
				newCar := Car{Name: name, Type: CarType(carType), CheckInTime: time.Now()}
				cp.suvCars[name] = newCar
				fmt.Println("HB parked in SUV section")
			} else {
				fmt.Println("Max occupancy reached")
			}
			return
		} else {
			newCar := Car{Name: name, Type: CarType(carType), CheckInTime: time.Now()}
			cp.hbCars[name] = newCar
			fmt.Println("HB car parked")
			return
		}
	}

}
func (cp *CarParking) getCount() {
	fmt.Printf("Number of HB cars is %d", len(cp.hbCars))
	fmt.Printf("Number of SUV cars is %d", len(cp.suvCars))
	return
}
func (cp *CarParking) getAmount(car Car, price int) int {
	tHours := time.Since(car.CheckInTime)
	amt := tHours.Hours() * float64(price)
	return int(amt)
}
func (cp *CarParking) exitCar(name string, carT CarType) {
	var totalAmt int
	if carT == CarType("suv") {
		if _, exists := cp.suvCars[name]; !exists {
			fmt.Println("This car doesn't exist")
			return
		}
		totalAmt = cp.getAmount(cp.suvCars[name], 20)
		delete(cp.suvCars, name)
	} else {
		if _, exists := cp.hbCars[name]; !exists {
			fmt.Println("This car doesn't exist")
			return
		}
		totalAmt = cp.getAmount(cp.hbCars[name], 10)
		delete(cp.hbCars, name)
	}
	fmt.Printf("Total amount is %d", totalAmt)
}
func (cp *CarParking) listParking() {
	fmt.Println(cp)
}
func ParkingLot() {
	newPL := NewCarParking()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		parts := strings.Fields(input)
		cmd, args := parts[0], parts[1:]

		switch cmd {
		case "park":
			newPL.parkVehicle(args[0], CarType(args[1]))
		case "exit":
			newPL.exitCar(args[0], CarType(args[1]))
		case "list":
			newPL.listParking()
		}
	}
}

type Vehicle struct {
	Vehicle uuid.UUID
	RegNo   string
}
type NewParkingLot struct {
	ParkingLot map[string]map[string]Vehicle
}

func MakeParkingLot(id string) *NewParkingLot {
	
	return &NewParkingLot{
		ParkingLot: make(map[string]map[string]Vehicle),
	}
}
