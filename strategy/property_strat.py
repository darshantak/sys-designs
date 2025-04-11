from abc import ABC, abstractmethod

class PropertyStrat(ABC):
    
    @abstractmethod
    def get_desired_property(self, properties):
        pass