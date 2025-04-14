from abc import ABC,abstractmethod

class Subscriber(ABC): #Subscriber
    @abstractmethod
    def update(self,message):
        pass

class Publisher(ABC): #Publisher
    @abstractmethod
    def add_observer(self,observer):
        pass
    
    @abstractmethod
    def remove_observer(self,observer):
        pass
    
    @abstractmethod
    def notify_all(self,message):
        pass
    
