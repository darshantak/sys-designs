import json
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
    
    def search_property(self,filters):
        results = self.properties
        
        if "price" in filters:
            obj = ByPriceRange(filters["price"])
            results = obj.get_desired_property(results)
                
        if "num_of_rooms" in filters:
            obj = ByRooms(filters["num_of_rooms"])
            results = obj.get_desired_property(results)
        
        return results
    
class StrategyFactory:
    def __init__(self,strategy):
        self.strategy =  strategy
    
    def get_strategy(self):
        if self.strategy == Strategy.PRICE_RANGE:
            return ByPriceRange
        else:
            return ByRooms
        