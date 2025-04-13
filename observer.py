from abc import ABC,abstractmethod

class Observer(ABC):
    @abstractmethod
    def update(self,message):
        pass

class Subject(ABC):
    @abstractmethod
    def add_observer(self,observer):
        pass
    
    @abstractmethod
    def remove_observer(self,observer):
        pass
    
    @abstractmethod
    def notify_all(self,message):
        pass
    
