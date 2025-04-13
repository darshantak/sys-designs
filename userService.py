from user import User
class UserService:
    def __init__(self) -> None:
        self.users = {}
        
    def registerUser(self, username, password):
        newUser = User(username, password)
        self.users[username] = newUser
        return newUser

    def loginUser(self, username, password):
        if username in self.users:
            if password == self.users[username].password:
                return True
        
        return False

