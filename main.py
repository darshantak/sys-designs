from parkingApp import ParkingApp
from vehicle import Vehicle
def main():
    app = ParkingApp()
    
    
    car = Vehicle("123","car")
    bike = Vehicle("ABC","bike")
    
    app.addParkingFloor("ABC123")
    app.addParkingFloor("PQR123")
    
    app.addParkingSlots("ABC123","SLOT1","car")
    app.addParkingSlots("ABC123","SLOT2","car")
    app.addParkingSlots("ABC123","SLOT3","car")
    
    app.parkVehicle(car)
    app.parkVehicle(bike)
    
    app.unParkVehicle("RJ")
    