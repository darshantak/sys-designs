from servicable import Servicable
class Answer(Servicable):
    def __init__(self,aid,author, content):
        self.aid = aid
        self.author = author
        self.content = content
        self.votes= 0
        self.comments = []
    
    def addComments(self,comment):
        self.comments.append(comment)
    
    def downVote(self):
        self.votes -= 1
        
    def upVote(self):
        self.votes += 1
        

    