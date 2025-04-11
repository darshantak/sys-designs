from strategy.property_strat import PropertyStrat

class ByRooms(PropertyStrat):
    def __init__(self,num_rooms) -> None:
        super().__init__()
        self.num_rooms= num_rooms
    
    def get_desired_property(self, properties):
        result_props = []
        for i in properties:
            if self.num_rooms <= i["rooms"]:
                result_props.append(i)
        
        return result_props
    
