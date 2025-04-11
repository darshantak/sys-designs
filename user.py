import json
class User:
    def __init__(self,name,contact):
        self.name = name
        self.contact = contact
        
    def __repr__(self):
        return json.dumps({
            "name": self.name,
            "contact": self.contact
        }, indent=4)