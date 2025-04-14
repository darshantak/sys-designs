from auctionService import AuctionService
from datetime import datetime,timedelta
def main():
    auction_system = AuctionService()
    auction_system.registerUser("alice", "alice123")
    auction_system.registerUser("bob", "bob123")

    # Step 3: Login Users (optional, for completeness)
    auction_system.loginUser("alice", "alice123")
    auction_system.loginUser("bob", "bob123")

    # Step 4: Create Auction Item
    end_time = datetime.now() + timedelta(seconds=10)  # auction ends in 10 sec
    auction_item = auction_system.createAuctionItem(
        aid=1,
        name="MacBook Pro",
        description="16-inch M1 Pro, 1TB SSD",
        startingPrice=1000,
        expirationTime=end_time
    )
    auction_item_1 = auction_system.createAuctionItem(
        aid=2,
        name="Windows",
        description="16-inch M1 Pro, 1TB SSD",
        startingPrice=1000,
        expirationTime=end_time
    )
    print(auction_item)
    
    # Step 5: Get user and item
    alice = auction_system.userService.users["alice"]
    bob = auction_system.userService.users["bob"]
    item = auction_system.auctionService.auctionItems[1]
    item_windows = auction_system.auctionService.auctionItems[2]
    print(alice)
    print(bob)
    print(item)
    
    # Step 6: Place Bids
    # item.add_observer(alice)
    # item.add_observer(bob)
    
    # Place a Bid – should trigger notifications
    bid1 = auction_system.createBid(bidId=1, item=item, user=alice, priceQuoted=1100)
    bid2 = auction_system.createBid(bidId=2, item=item, user=bob, priceQuoted=1200)
    bid3 = auction_system.createBid(bidId=2, item=item_windows, user=alice, priceQuoted=1200)
    bid4 = auction_system.createBid(bidId=2, item=item_windows, user=bob, priceQuoted=1500)
    
    print(bid1)
    print(bid2)
    auction_system.getAllBids()
    
    
    
main()