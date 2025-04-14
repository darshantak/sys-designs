from socialMediaApp import SocialMediaApp
from strategy import sortByLikes
def main():
    app = SocialMediaApp()
    
    print("✅ Creating user accounts...")
    alice = app.createAccount("Alice", "alice@example.com", "alice123")
    bob = app.createAccount("Bob", "bob@example.com", "bob123")
    charlie = app.createAccount("Charlie", "charlie@example.com", "charlie123")
    print(f"👤 Created: {alice}")
    print(f"👤 Created: {bob}")
    print(f"👤 Created: {charlie}")
    
    # 3. Valid Login
    print("\n✅ Logging in users...")
    loggedInAlice = app.loginUser("alice@example.com", "alice123")
    loggedInBob = app.loginUser("bob@example.com", "bob123")
    loggedInCharlie = app.loginUser("charlie@example.com", "charlie123")
    print(f"🔐 Alice: {loggedInAlice}")
    print(f"🔐 Bob: {loggedInBob}")
    print(f"🔐 Charlie: {loggedInCharlie}")

    # 4. Profile Setup
    print("\n📝 Setting up user profiles...")
    profileAlice = app.createProfile(loggedInAlice, "Artist and Music lover.", ["art", "music", "design"], "alice_pic.jpg")
    profileBob = app.createProfile(loggedInBob, "Engineer. Gamer. Thinker.", ["coding", "gaming"], "bob_pic.png")
    profileCharlie = app.createProfile(loggedInCharlie, "Explorer of life!", ["travel", "food", "books"], "charlie_pic.bmp")
    print(f"🧬 Alice Profile: {profileAlice}")
    print(f"🧬 Bob Profile: {profileBob}")
    print(f"🧬 Charlie Profile: {profileCharlie}")

    # 5. Connection Requests
    print("\n🔗 Sending connection requests...")
    print("📨 Alice → Bob:", app.sendConnectionRequest(loggedInAlice, loggedInBob))
    print("📨 Bob → Charlie:", app.sendConnectionRequest(loggedInBob, loggedInCharlie))
    print("📨 Charlie → Alice:", app.sendConnectionRequest(loggedInCharlie, loggedInAlice))

    # 6. Accepting Connection Requests
    print("\n🤝 Accepting connection requests...")
    print("✅ Bob accepts Alice:", app.acceptConnectionRequest(loggedInBob, loggedInAlice))
    print("✅ Charlie accepts Bob:", app.acceptConnectionRequest(loggedInCharlie, loggedInBob))
    print("✅ Alice accepts Charlie:", app.acceptConnectionRequest(loggedInAlice, loggedInCharlie))
    
    # 8. Listing Connections
    print("\n📋 Connections List:")
    print("Alice's friends:", app.listConnections(loggedInAlice))
    print("Bob's friends:", app.listConnections(loggedInBob))
    print("Charlie's friends:", app.listConnections(loggedInCharlie))

    # 9. Post Creation
    print("\n📝 Creating posts...")
    post1 = app.createPost(loggedInAlice, "Beautiful sunset today!", "sunset.jpg")
    post2 = app.createPost(loggedInBob, "Just built a gaming PC 😍", "rig.png")
    post3 = app.createPost(loggedInCharlie, "Trying a new recipe today.", "cooking.jpg")
    print(f"📮 Post 1: {post1}")
    print(f"📮 Post 2: {post2}")
    print(f"📮 Post 3: {post3}")

    
    # 10. News Feed for Each User
    print("\n📰 Loading News Feed for Alice:")
    app.setNewStrategy(sortByLikes.SortByLikes())
    app.loadNewsFeed(loggedInAlice)

    print("\n📰 Loading News Feed for Bob:")
    app.loadNewsFeed(loggedInBob)

    print("\n📰 Loading News Feed for Charlie:")
    app.loadNewsFeed(loggedInCharlie)

    # 11. Isolated User Edge Case
    print("\n🧍‍♂️ Creating a new isolated user Dave (no connections)...")
    dave = app.createAccount("Dave", "dave@example.com", "dave123")
    loggedInDave = app.loginUser("dave@example.com", "dave123")
    profileDave = app.createProfile(loggedInDave, "Just joined!", ["chess"], "dave.jpg")
    print("📰 Dave's empty feed:")
    app.loadNewsFeed(loggedInDave)

    print("\n✅ Social Media App Test Completed\n")
    
    
    
        
main()
    