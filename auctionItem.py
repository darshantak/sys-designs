class AuctionItem:
    def __init__(self, aid, name, description, startingPrice,expirationTime) -> None:
        self.aid = aid
        self.name = name
        self.description = description
        self.startingPrice = startingPrice
        self.expirationTime = expirationTime
        self.bids = []
        
    def __str__(self):
        # return f"AuctionItem(id={self.aid}, name={self.name}, price={self.startingPrice}, ends={self.expirationTime})"
        return (f"AuctionItem("f"expires={self.expirationTime}" f"bids={self.bids})")

    def __repr__(self):
        return self.__str__()
    