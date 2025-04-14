from userService import UserService
from userProfileService import UserProfileService
from connectionService import ConnectionService
from postService import PostService
from newsFeedService import NewsFeedService
class SocialMediaApp:
    def __init__(self) -> None:
        self.userService = UserService()
        self.profileService = UserProfileService()
        self.connectionService = ConnectionService()
        self.postService = PostService()
        self.newsFeedService = NewsFeedService(self.postService,self.connectionService)
            
    def createAccount(self, name: str, email: str, password: str):
        newUser = self.userService.createAccount(name, email, password)
        return newUser
    
    def loginUser(self, email: str, password: str):
        return self.userService.loginUser(email,password)
    
    def createProfile(self, user, bio: str, interests: list[str], profilePicture: str):
        newProfile = self.profileService.createProfile(user,bio,interests,profilePicture)
        return newProfile
    
    def createPost(self, user, description: str, image: str):
        newPost = self.postService.createPost(user,description,image)
        return newPost
    
    def likePost(self, user, post):
        self.postService.likePost(user, post)
        
    def sendConnectionRequest(self, byUser, toUser):
        return self.connectionService.sendConnectionRequest(byUser, toUser)
    
    def acceptConnectionRequest(self, user, toAcceptUser):
        return self.connectionService.acceptConnectionRequest(user, toAcceptUser)

    def listConnections(self, user):
        list = self.connectionService.listConnections(user)
        return list
    
    def setNewStrategy(self, strategy):
        self.newsFeedService.setNewStrategy(strategy)
        
    def loadNewsFeed(self,user):
        posts = self.newsFeedService.loadNewsFeed(user)
        for post in posts:
            print(post) 
               
        
        
    