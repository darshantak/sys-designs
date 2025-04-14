from strategy.strategy import Strategy

class SortByTime(Strategy):
    def sort_all(self,posts):
        return sorted(posts,key = lambda post:post.createdOn, reverse= True)
    
