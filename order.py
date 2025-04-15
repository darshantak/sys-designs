import uuid

class Order:
    def __init__(self, orderId: str, orderType, user, book):
        self.orderId = orderId
        self.orderType = orderType
        self.user = user
        self.book = book

    def __repr__(self):
        return f"<Order {self.orderId} | {self.orderType.value} | User: {self.user.name} | Book: {self.book.title}>"

