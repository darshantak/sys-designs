package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type VehicleType string

const (
	Car  VehicleType = "car"
	Bike VehicleType = "bike"
)

type Vehicle struct {
	Type   VehicleType
	RegNum string
}
type ParkingSlot struct {
	SlotNumber int
	SlotType   VehicleType
	IsOccupied bool
	Vehicle    *Vehicle
}
type ParkingFloor struct {
	FloorNumber int
	Slots       map[int]*ParkingSlot
}
type ParkingLot struct {
	Id            string
	ParkingFloors map[int]*ParkingFloor
}

type Ticket struct {
	TicketId string
	Vehicle  *Vehicle
}

func NewParkingLot(id string) *ParkingLot {
	return &ParkingLot{
		Id:            id,
		ParkingFloors: make(map[int]*ParkingFloor),
	}
}

func (pl *ParkingLot) AddFloor(floor *ParkingFloor) {
	pl.ParkingFloors[floor.FloorNumber] = floor
}

func (pf *ParkingFloor) AddSlot(slot *ParkingSlot) {
	pf.Slots[slot.SlotNumber] = slot
}

func (pl *ParkingLot) FindAvailableSlot(vehicleType VehicleType) (*ParkingFloor, *ParkingSlot, error) {
	for _, i := range pl.ParkingFloors {
		for _, j := range i.Slots {
			if !j.IsOccupied && vehicleType == j.SlotType {
				return i, j, nil
			}
		}
	}
	return nil, nil, errors.New("no available slot")
}

func (pl *ParkingLot) ParkVehicle(vehicle Vehicle) (*Ticket, error) {
	floor, slot, err := pl.FindAvailableSlot(vehicle.Type)
	if err != nil {
		return nil, err
	}
	ticketId := fmt.Sprintf("%s_%d_%d", pl.Id, floor.FloorNumber, slot.SlotNumber)
	ticket := Ticket{
		TicketId: ticketId,
		Vehicle:  &vehicle,
	}
	return &ticket, nil
}

func (pl *ParkingLot) UnParkVehicle(ticketId string) error {
	// var floorNumber, slotNumber string
	splits := strings.Split(ticketId, "_")
	floorNumber, err := strconv.Atoi(splits[1])
	if err != nil {
		return err
	}
	slotNumber, err := strconv.Atoi(splits[2])
	if err != nil {
		return err
	}
	floor, exists := pl.ParkingFloors[floorNumber]
	if !exists {
		return errors.New("floor does not exist")
	}

	slot, exists := floor.Slots[slotNumber]
	if !exists {
		return errors.New("slot does not exist")
	}

	if slot.IsOccupied {
		return errors.New("slots already empty")
	}

	slot.IsOccupied = false
	slot.Vehicle = nil

	return nil
}

func main() {
	parkingLot := NewParkingLot("PR1234")
	for i := 1; i <= 3; i++ {
		floor := &ParkingFloor{FloorNumber: i, Slots: make(map[int]*ParkingSlot)}
		parkingLot.AddFloor(floor)
		floor.AddSlot(&ParkingSlot{SlotNumber: 1, SlotType: Bike})
		floor.AddSlot(&ParkingSlot{SlotNumber: 2, SlotType: Bike})
		for j := 3; j <= 7; j++ {
			floor.AddSlot(&ParkingSlot{SlotNumber: j, SlotType: Car})
		}
	}

	vehicle1 := &Vehicle{Type: Car, RegNum: "KA-01-HH-1234"}

	ve1, err := parkingLot.ParkVehicle(*vehicle1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ve1.TicketId)
}
