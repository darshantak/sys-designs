from abc import ABC,abstractmethod

class Servicable(ABC):
    
    @abstractmethod
    def addComments(self):
        pass

    def upVote(self):
        pass
    
    def downVote(self):
        pass
    
    