from user import User

class UserService:
    def __init__(self):
        self.users = {} 

    def createAccount(self, name: str, email: str, password: str) -> User:
        if name in self.users:
            raise ValueError("User already exists with this name.")
        user = User(name, email, password)
        self.users[name] = user
        return user

    def loginUser(self, email: str, password: str) -> User:
        for user in self.users.values():
            if user.email == email and user.password == password:
                return user
        raise ValueError("Invalid email or password.")