from strategy import sortByTime
class NewsFeedService:
    def __init__(self, postService, connectionService):
        self.postService = postService
        self.connectionService = connectionService
        self.strategy = sortByTime.SortByTime()
        
    def getConnectionsOfUser(self, user) -> list:
        return self.connectionService.listConnections(user)

    def getPostsOfUser(self, user) -> list:
        return self.postService.listPostByUser(user)

    def setNewStrategy(self,strategy):
        self.strategy = strategy
        
    def loadNewsFeed(self, user) -> list:
        connections = self.getConnectionsOfUser(user)
        all_posts = []

        for connection in connections:
            posts = self.getPostsOfUser(connection)
            all_posts.extend(posts)

        # sorted_posts = sorted(all_posts, key=lambda post: post.createdOn, reverse=True)
        sorted_posts = self.strategy.sort_all(all_posts)
        return sorted_posts
