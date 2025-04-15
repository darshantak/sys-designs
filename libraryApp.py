from bookService import BookService
from memberService import MemberService
from orderService import OrderService

class LibraryApp:
    def __init__(self):
        self.bookService = BookService()
        self.memberService = MemberService()
        self.orderService = OrderService()
        
    def addBook(self, title, author, isbn, publication):
        return self.bookService.addBook(title, author, isbn, publication)

    def updateBook(self, isbn, **kwargs):
        book = self.bookService.books.get(isbn)
        if not book:
            print(f"[WARN] Book with ISBN {isbn} not found.")
            return None
        return self.bookService.updateBook(book, **kwargs)
    
    def deleteBook(self, isbn):
        book = self.bookService.books.get(isbn)
        if not book:
            print(f"[WARN] Book with ISBN {isbn} does not exist.")
            return False
        return self.bookService.deleteBook(book)

    def registerMember(self, name, memberID, contact):
        return self.memberService.registerMember(name, memberID, contact)

    def borrowBook(self, user, book):
        return self.orderService.borrowBook(user, book)

    def returnBook(self, user, book):
        return self.orderService.returnBook(user, book)
    
    