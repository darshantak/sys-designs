from bid import Bid

class BidService:
    def __init__(self) -> None:
        self.bids= {}
    
    def createBid(self,bidId , item , user , priceQuoted ):
        newBid = Bid(bidId, item, user=user, priceQuoted=priceQuoted)
        self.bids[bidId] = newBid
        return newBid
    
    