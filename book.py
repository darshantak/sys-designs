class Book:
    def __init__(self, title: str, author: str, isbn: str, publication: str, available: bool = True):
        self.title = title
        self.author = author
        self.isbn = isbn
        self.publication = publication
        self.available = available

    def __repr__(self):
        return f"<Book {self.title} by {self.author} | ISBN: {self.isbn} | Available: {self.available}>"
