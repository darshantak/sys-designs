from auctionItem import AuctionItem
from bid import Bid

class AuctionItemService:
    def __init__(self) -> None:
        self.auctionItems = {}
        
    def createAuctionItem(self,aid, name, description, startingPrice,expirationTime):
        newItem = AuctionItem(aid, name, description, startingPrice,expirationTime)
        self.auctionItems[aid] = newItem
        return newItem
    
    def searchAuctionItem(self,name):
        pass
    
    def browseAuctionItems(self):
        return self.auctionItems
    
    def addBids(self,auctionItem:AuctionItem,bid:Bid):
        # itemId = auctionItem.aid
        auctionItem.addBids(bid)
        # auctionItem.notify_all(
        # f"New bid of {bid.priceQuoted} placed on '{auctionItem.name}' by {bid.user.username}"
        # )

