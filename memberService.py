from member import Member

class MemberService:
    def __init__(self):
        self.members = {}  # map from memberID to Member object

    def registerMember(self, name: str, memberID: str, contact: str):
        if memberID in self.members:
            print(f"[WARN] Member ID {memberID} already exists.")
            return None
        new_member = Member(name, memberID, contact)
        self.members[memberID] = new_member
        print(f"[INFO] Registered new member: {new_member}")
        return new_member
    
    def getMember(self, memberID: str):
        return self.members.get(memberID, None)
    