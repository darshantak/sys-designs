import json
from enums import property_type
class Property():
    def __init__(self,user,type,house_address,location,num_of_rooms,price,security_deposit,deal_type):
        self.user = user
        self.type = type
        self.house_address = house_address
        self.location =location
        self.num_of_rooms = num_of_rooms
        self.price = price
        self.security_deposit = security_deposit
        self.deal_type = deal_type
    

    def __repr__(self):
        return json.dumps({
            "user": self.user,
            "house_address": self.house_address,
            "type" : self.type,
            "deal_type": self.deal_type,
            "price":self.price
        }, indent=4,default=str)
    
    