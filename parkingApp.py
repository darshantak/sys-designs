from parkingFloorService import ParkingFloorService
from slotService import SlotService
class ParkingApp:
    def __init__(self) -> None:
        self.parkingFloorService = ParkingFloorService()
        self.slotService = SlotService()
        self.globalLot = {}
        
    def backfiller(self):
        floors = self.parkingFloorService.listFloors()
        # slots = self.slotService.
        
    def addParkingFloor(self,floorID):
        self.parkingFloorService.addParkingFloor(floorID,capacity)
        return True
    
    def addParkingSlots(self,floorId,slotId,type):
        if self.parkingFloorService.isExist(floorId):
            self.slotService.addSlot(floorId, slotId, type)
        else:
            return False
    
    def parkVehicle(self, vehicle):
        pass

    def unParkVehicle(self,vehicle):
        pass

    