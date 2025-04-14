from post import Post
class PostService:
    def __init__(self):
        self.posts = {}  # Dictionary mapping User to list of Post objects

    def createPost(self, user, description: str, image: str) -> Post:
        post = Post(description, image)
        if user not in self.posts:
            self.posts[user] = []
        self.posts[user].append(post)
        return post

    def likePost(self, user, post: Post) -> bool:
        if post:
            post.likes += 1
            return True
        return False

    def commentPost(self, user, post: Post, comment: str) -> bool:
        if post:
            post.comments.append(comment)
            return True
        return False

    def listPostByUser(self, user) -> list:
        return self.posts.get(user, [])
