from strategy.property_strat import PropertyStrat

class ByPriceRange(PropertyStrat):

    def __init__(self,range) -> None:
        super().__init__()
        self.low_range = range[0]
        self.high_range = range[1]
        
    
    def get_desired_property(self,properties):
        result_props = []
        print(properties["User 1"]["user"])
        for i in properties:
            
            if self.low_range <= v["price"] and v["price"]<=self.high_range:
                result_props.append(i)
        return result_props
    