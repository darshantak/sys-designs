from parkingFloor import ParkingFloor
class ParkingFloorService:
    def __init__(self) -> None:
        self.floors = {}

    def addParkingFloor(self,floorID):
        newFloor = ParkingFloor(floorID,capacity)
        self.floors[floorID] = newFloor
        return newFloor

    def updateFloor(self, floorId, slot):
        pass
