# Parking Lot System

## Functionalities

The parking lot system can perform the following functions:

- **Create the parking lot.**
- **Add floors to the parking lot.**
- **Add a parking slot to any of the floors.**
- **Park a vehicle:** 
    - Finds the first available slot for the given vehicle type.
    - Books the slot, creates a ticket, parks the vehicle, and returns the ticket.
- **Unpark a vehicle** given the ticket ID.
- **Display the number of free slots per floor** for a specific vehicle type.
- **Display all free slots per floor** for a specific vehicle type.
- **Display all occupied slots per floor** for a specific vehicle type.

## Vehicle Details

Every vehicle will have:
- **Type**
- **Registration number**
- **Color**

### Vehicle Types

The system will handle the following types of vehicles:

- Car
- Bike
- Truck

## Parking Slot Details

Each parking slot will have the following constraints:

- Each slot is designated for a specific type of vehicle. Only the corresponding vehicle type should be allowed in that slot.
- **Finding the first available slot** for a vehicle must follow these rules:
    - The slot type should match the vehicle type.
    - The slot should be on the **lowest possible floor** in the parking lot.
    - The slot should have the **lowest possible slot number** on that floor.
- Slots are numbered serially from **1 to n** for each floor, where **n** is the number of parking slots on that floor.

## Parking Lot Floor Details

- Floors are numbered serially from **1 to n**, where **n** is the total number of floors.
- Each floor may contain one or more parking slots of different types.
- Slot allocation on each floor follows a specific order:
    - The first slot is for **trucks**.
    - The next two slots are for **bikes**.
    - All remaining slots are for **cars**.

## Ticket Details

- Each ticket ID follows the format:
  <parking_lot_id><floor_no><slot_no>


- **Example:** `PR1234_2_5` (indicates the 5th slot on the 2nd floor of parking lot `PR1234`).

- The parking lot ID is **PR1234** (we assume there is only one parking lot).


