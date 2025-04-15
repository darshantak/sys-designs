import uuid
from order import Order
class OrderService:
    
    def __init__(self) -> None:
        self.orders = {}
        self.usersToBooks = {}
    
    def borrowBook(self, user, book):
        
        if not book.available:
            print(f"[WARN] Book '{book.title}' is not available for borrowing.")
            return None

        book.availability = False
        orderId = str(uuid.uuid4())
        newOrder = Order(orderId,"BORROW",book,user)
        self.orders.setdefault(user,[]).append(newOrder)
        self.usersToBooks.setdefault(user,[]).append(book)
        print(f"[INFO] {user.name} borrowed '{book.title}'.")

        return newOrder

    def returnBook(self, user, book):
        if book not in self.usersToBooks.get(user,[]):
            print(f"[WARN] Book '{book.title}' is not borrowed by {user.name}.")
            return None

        orderId = str(uuid.uuid4())
        order = Order(orderId,"RETURN", user, book)

        book.available = True
        self.orders.setdefault(user, []).append(order)
        self.usersToBooks[user].remove(book)
        
        print(f"[INFO] {user.name} returned '{book.title}'.")
        return order