from connection import Connection
class ConnectionService:
    def __init__(self):
        self.connections = {}  # Dictionary mapping User to Connection

    def _get_or_create_connection(self, user):
        if user not in self.connections:
            self.connections[user] = Connection()
        return self.connections[user]

    def sendConnectionRequest(self, byUser, toUser) -> bool:
        if byUser == toUser:
            return False

        toUserConnection = self._get_or_create_connection(toUser)
        if byUser in toUserConnection.connectionsList or byUser in toUserConnection.pendingConnectionList:
            return False

        toUserConnection.pendingConnectionList.append(byUser)
        return True

    def acceptConnectionRequest(self, user, toAcceptUser) -> bool:
        userConnection = self._get_or_create_connection(user)
        if toAcceptUser not in userConnection.pendingConnectionList:
            return False

        userConnection.pendingConnectionList.remove(toAcceptUser)
        userConnection.connectionsList.append(toAcceptUser)

        toAcceptConnection = self._get_or_create_connection(toAcceptUser)
        toAcceptConnection.connectionsList.append(user)
        return True

    def declineConnectionRequest(self, user, toAcceptUser) -> bool:
        userConnection = self._get_or_create_connection(user)
        if toAcceptUser in userConnection.pendingConnectionList:
            userConnection.pendingConnectionList.remove(toAcceptUser)
            return True
        return False

    def listConnections(self, user):
        userConnection = self._get_or_create_connection(user)
        return userConnection.connectionsList
