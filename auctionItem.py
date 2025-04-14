from observer import Publisher

class AuctionItem(Publisher):
    def __init__(self, aid, name, description, startingPrice,expirationTime) -> None:
        self.aid = aid
        self.name = name
        self.description = description
        self.startingPrice = startingPrice
        self.expirationTime = expirationTime
        self.bids = []
        self.observers = set()
        
    def add_observer(self, observer):
        self.observers.add(observer)
        
    def remove_observer(self, observer):
        self.observers.discard(observer)
        
    def notify_all(self, message):
        for observer in self.observers:
            observer.update(message)

    def addBids(self,bid):
        self.bids.append(bid)
        self.add_observer(bid)
        self.notify_all(f"New bid of {bid.priceQuoted} placed on '{self.name}' by {bid.user}")
        
    def __str__(self):
        return (f"AuctionItem("f"expires={self.expirationTime} id={self.aid}, name={self.name}" f"bids={self.bids})")

    def __repr__(self):
        return self.__str__()
    