class Member:
    def __init__(self, name: str, memberID: str, contact: str):
        self.name = name
        self.memberID = memberID
        self.contact = contact

    def __repr__(self):
        return f"<Member {self.name} | ID: {self.memberID} | Contact: {self.contact}>"
