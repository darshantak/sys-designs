from datetime import datetime
class Bid:
    def __init__(self,bidId,item,user,priceQuoted) -> None:
        self.bidId = bidId
        self.item = item
        self.user = user
        self.priceQuoted = priceQuoted
        
    def __str__(self):
        return f"Bid(id={self.bidId}, user={self.user.username}, item={self.item.name}, price={self.priceQuoted})"

    def __repr__(self):
        return self.__str__()