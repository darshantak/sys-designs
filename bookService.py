from book import Book

class BookService:
    def __init__(self) -> None:
        self.books={}
    
    def addBook(self, title, author, isbn, publication):
        if isbn in self.books:
            print(f"[INFO] Book with ISBN {isbn} already exists.")
            return None
        new_book = Book(title, author, isbn, publication)
        self.books[isbn] = new_book
        return new_book
    
    def updateBook(self, isbn, **kwargs):
        book = self.books.get(isbn)
        if not book:
            return None
        for key, value in kwargs.items():
            if hasattr(book, key):
                setattr(book, key, value)
        return book

    def deleteBook(self,isbn):
        if isbn in self.books:
            del self.books[isbn]
            return True
    
        return False
    
    def getAvalaibiltyOfBook(self,isbn):
        if isbn in self.books:
            book = self.books[isbn]
            return book.available
        
        return False
    
        