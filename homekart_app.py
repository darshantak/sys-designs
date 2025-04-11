from user_service import UserService
from property_service import PropertyService
from booking_service import BookingService
class HomeKartApp:
    def __init__(self) -> None:
        self.user_service = UserService()
        self.property_service= PropertyService()
        self.booking_service = BookingService()
    
    def add_user(self,name,contact):
        print("\n\n------------------- Add User -------------------\n\n")
        res = self.user_service.add_user(name,contact)
        print(res)
        
    def add_property(self,user,type,house_address,location,num_of_rooms,price,security_deposit,deal_type):
        print("\n\n------------------- Add Property -------------------\n\n")
        res = self.property_service.add_property(user,type,house_address,location,num_of_rooms,price,security_deposit,deal_type)
        print(res)
    
    def search_property(self,filters):
        print("\n\n------------------- Search Properties -------------------\n\n")
        res = self.property_service.search_property(filters)
        for i in res:
            print(res)
    
        
        
        
