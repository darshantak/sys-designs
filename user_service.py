from user import User
class UserService:
    def __init__(self) -> None:
        self.users = {}
        
    def add_user(self,name,contact):
        if name in self.users:
            raise Exception("User already exists with same name")
        user = User(name,contact)
        self.users[name] = user    
        return user
    
    
