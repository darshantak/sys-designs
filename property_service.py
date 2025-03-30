from property import Property
from enums.strategy import Strategy
from strategy.by_num_rooms import ByRooms
from strategy.by_price_range import ByPriceRange

class PropertyService:
    def __init__(self) -> None:
        self.properties = {}
    
    def add_property(self,user,type,house_address,location,num_of_rooms,price,security_deposit,deal_type):
        property = Property(user,type,house_address,location,num_of_rooms,price,security_deposit,deal_type)
        self.properties[user] = property
        return property
    
    def search_property(self,strategy):
        pass
    

class StretegyFactory:
    def __init__(self,strategy):
        self.strategy =  strategy
    
    def get_strategy(self):
        if self.strategy == Strategy.PRICE_RANGE:
            return ByPriceRange
        else:
            return ByRooms
        