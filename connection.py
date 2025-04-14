class Connection:
    def __init__(self):
        self.connectionsList = []      # List of accepted User connections
        self.pendingConnectionList = []  # List of incoming connection requests (Users)

    def __str__(self):
        return f"ConnectionsList:{self.connectionsList} PendingConnectionList:{self.pendingConnectionList}"