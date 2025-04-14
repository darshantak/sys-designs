from userProfile import UserProfile
class UserProfileService:
    def __init__(self):
        self.profiles = {}  # Dictionary mapping User objects to Profile objects

    def createProfile(self, user, bio: str, interests: list[str], profilePicture: str) -> UserProfile:
        if user in self.profiles:
            raise ValueError("Profile already exists for this user.")
        profile = UserProfile(bio, interests, profilePicture)
        self.profiles[user] = profile
        return profile

    def updateProfile(self, user, **kwargs) -> bool:
        if user not in self.profiles:
            raise ValueError("Profile does not exist for this user.")
        
        profile = self.profiles[user]

        for key, value in kwargs.items():
            if hasattr(profile, key):
                setattr(profile, key, value)
        return True
