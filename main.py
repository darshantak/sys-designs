from stackoverflow import StackOverFlowApp
def main():
    app = StackOverFlowApp()

    # Add some users
    print(app.addUser(1))
    print(app.addUser(2))

    # # Add a question by user1
    # question = app.addQuestion(qid=101, authorId=1, content="What is Python?")
    # print(f"Question added: {question.content}")
    # # Add an answer to the question by user2
    # answer = app.addAnswer(qid=101, aid=201, authorId=2, content="Python is a powerful programming language.")
    # print(f"Answer added: {answer.content}")
    # # Add a comment to the question
    # comment1 = app.addComment(target=question, cid=301, authorId=2, content="Great question!")

    # # Add a comment to the answer
    # comment2 = app.addComment(target=answer, cid=302, authorId=1, content="Thanks for the answer!")

    # # Upvote the question and answer
    # app.upVote(question)
    # app.upVote(answer)

    # # Downvote the question
    # app.downVote(question)

    # # Print out the votes to verify
    # print(f"Votes on Question: {question.votes}")
    # print(f"Votes on Answer: {answer.votes}")

    # # Optionally: print the comments to verify
    # print("Question Comments:")
    # for c in question.comments:
    #     print(f"- {c.content} by user {c.authorId}")

    # print("Answer Comments:")
    # for c in answer.comments:
    #     print(f"- {c.content} by user {c.authorId}")
        
main()