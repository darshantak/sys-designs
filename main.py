from homekart_app import HomeKartApp
from enums.property_type import PropertyType
from enums.deal_type import DealType
def main():
    app = HomeKartApp()
    app.add_user("User 1", 123344556)
    app.add_user("User 2", 987654545)
    
    app.add_property("User 1", PropertyType.FLAT,"#102, Mulberry Coast", 560023, 2, 50000,250000,DealType.RENT)
    app.add_property("User 2", PropertyType.INDEPENDENT,"#602, Adarsh Palm", 560023, 1, 1500000,2500000,DealType.RENT)
    
    app.add_property("User 3", PropertyType.FLAT,"#102, Mulberry Coast", 560023, 6, 43500,250000,DealType.RENT)
    app.add_property("User 4", PropertyType.INDEPENDENT,"#602, Adarsh Palm", 560023, 4, 500000,2500000,DealType.RENT)
    
    app.add_property("User 5", PropertyType.FLAT,"#102, Mulberry Coast", 560023, 3, 500000,250000,DealType.RENT)
    app.add_property("User 6", PropertyType.INDEPENDENT,"#602, Adarsh Palm", 560023, 5, 90000,2500000,DealType.RENT)
    
    app.search_property(filters={"price":(0, 100000)})
    app.search_property(filters={"price":(0, 100000),"num_of_rooms":4})
    
if __name__ == "__main__":
    main()
    
    
    
    
    
    
    
    
    
