from datetime import datetime

class Post:
    def __init__(self, description: str, image: str):
        self.description = description
        self.image = image
        self.likes = 0
        self.comments = []
        self.createdOn = datetime.now()

    def __str__(self):
        return f"Post(description:{self.description},CreatedOn:{self.createdOn})"
    
    