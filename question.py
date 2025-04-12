from servicable import Servicable
class Question(Servicable):
    def __init__(self,qid,author,content) -> None:
        self.qid = qid
        self.author = author
        self.content = content
        self.votes = 0
        self.comments = []
        self.answers = []
        self.tags = []
        
    def addAnswer(self,answer):
        self.answers.append(answer)

    def addComments(self,comment):
        self.comments.append(comment)
    
    def downVote(self):
        self.votes -= 1
    
    def upVote(self):
        self.votes += 1
