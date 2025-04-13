from bidService import BidService
from auctionItemService import AuctionItemService
from userService import UserService
from bid import Bid

class AuctionService:
    def __init__(self) -> None:
        self.userService = UserService()
        self.auctionService = AuctionItemService()
        self.bidService = BidService()

    def registerUser(self, username, password):
        self.userService.registerUser(username,password)
        
    def loginUser(self,username, password):
        self.userService.loginUser(username=username, password=password)
        
    def createAuctionItem(self,aid, name, description, startingPrice,expirationTime):
        return self.auctionService.createAuctionItem(aid, name, description, startingPrice, expirationTime)

        
    def createBid(self,bidId , item , user , priceQuoted):
        newBid = Bid(bidId, item, user=user, priceQuoted=priceQuoted)
        self.auctionService.updateBids(item,newBid)
        return newBid

    def getAllBids(self):
        print(self.auctionService.auctionItems)
    