class UserProfile:
    def __init__(self, bio: str, interests: list[str], profilePicture: bytes):
        self.bio = bio
        self.profilePicture = profilePicture
        self.interests = interests

    def __str__(self):
        return f"UserProfile(bio:{self.bio}, interests:{self.interests})"
    
    # def __repr__(self) -> str:
    #     return f"UserProfile(bio:{self.bio}r!, interests:{self.interests})"
    