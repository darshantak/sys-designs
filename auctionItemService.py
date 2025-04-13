from auctionItem import AuctionItem

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
    
    def updateBids(self,auctionItem:AuctionItem,bid):
        itemId = auctionItem.aid
        self.auctionItems[itemId].bids.append(bid)

