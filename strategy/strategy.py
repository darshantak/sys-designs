from abc import abstractmethod,ABC

class Strategy(ABC):
    @abstractmethod
    def sort_all(self,posts):
        pass
    
