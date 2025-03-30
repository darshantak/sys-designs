from abc import ABC, abstractmethod

class PropertyStrat(ABC):
    
    @abstractmethod
    def get_desired_strategy(self, properties):
        pass