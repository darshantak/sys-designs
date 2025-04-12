from user import User
from question import Question
from answer import Answer
from comment import Comment
class StackOverFlowApp:
    def __init__(self) -> None:
        self.users = {}
        self.questions = {}
        
    def addUser(self,userId):
        newUser = User(userId)
        self.users[userId] = newUser
        return newUser
    
    def addQuestion(self,qid,authorId,content):
        newQuestion = Question(qid,authorId,content)
        self.questions[qid] = newQuestion
        self.users[authorId].questions.append(newQuestion)
        return newQuestion  
    
    def addAnswer(self,qid,aid,authorId,content):
        newAnswer = Answer(aid,authorId,content)
        if qid in self.questions:
            res = self.questions[qid]
            res.answers.append(newAnswer)
            self.users[authorId].answers.append(newAnswer)
            return newAnswer
        else:
            return False
    
    def addComment(self,target, cid, authorId, content):
        newComment = Comment(cid,authorId,content)
        target.addComments(newComment)
        return newComment
    
    def upVote(self,target):
        target.upVote()
    
    def downVote(self,target):
        target.downVote()
    
    
    def searchQuestion(self,tags):
        pass
    