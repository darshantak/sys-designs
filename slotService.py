from slot import Slot


class SlotService:
    def __init__(self) -> None:
        self.slots= {}
    
    def addSlot(self,floorId, slotId, type):
        newSlot = Slot(floorId, slotId, type, false)
        self.slots[slotId] = newSlot
        floors = self.parkingFloorService.listFloors()
        for i in floors:
            if i==floorId:
                i.slots.append(newSlot)
                
        return newSlot
