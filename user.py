class User:
	def __init__(self, username, password):
		self.username = username
		self.password = password
		
	def get_username(self):
		return self.username

	def get_password(self):
		return self.password
	
	def __str__(self):
		return f"User(username={self.username})"

	def __repr__(self):
		return self.__str__()