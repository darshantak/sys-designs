from libraryApp import LibraryApp
def main():
    app = LibraryApp()

    # Add books
    book1 = app.addBook("The Alchemist", "Paulo Coelho", "ISBN101", "HarperOne")
    book2 = app.addBook("1984", "George Orwell", "ISBN102", "Penguin")
    print(book1)
    print(book2)
    # Register members
    alice = app.registerMember("Alice", "M001", "alice@example.com")
    bob = app.registerMember("Bob", "M002", "bob@example.com")
    print(alice)
    print(bob)
    # Alice borrows a book
    app.borrowBook(alice, book1)

    # Bob tries to borrow the same book (should warn)
    app.borrowBook(bob, book1)

    # Alice returns the book
    app.returnBook(alice, book1)

    # Bob successfully borrows it now
    app.borrowBook(bob, book1)

    # Update book details
    app.updateBook("ISBN102", title="Nineteen Eighty-Four", publication="Updated Penguin")

    # Delete a book
    app.deleteBook("ISBN101")

if __name__ == "__main__":
    main()
