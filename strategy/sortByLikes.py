from strategy.strategy import Strategy
class SortByLikes(Strategy):
    def sort_all(self, posts:list):
        return sorted(posts,key=lambda post:post.likes, reverse=True)
    