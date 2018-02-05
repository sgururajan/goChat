# GoChat server

A simple chat server written in golang. It uses websocket to message tranfer and REST api for other data transfer.

# Build and Run

Run go run goChatServer.go

The server will run by default in port 5020. To change the port run with --port <port number>. The server also uses in memory database by default and can run with MongoDB.
To change the server to use Mongo run with "--inmemory false" option.

# Features

1. Authentication with JWT token
2. Web socket endpoint for message transfer
3. REST API endpoints for conversations, list of messages in a conversation, login and logout
4. In memory and MongoDB database support 