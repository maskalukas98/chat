## ABOUT

Trying out Golang + Sample Project.

User service for creating and retrieving users using REST and gRPC.

Sample tests only in ./user service

The user connects to the chat service using WebSocket, sends a new message to another user, which is stored in MongoDB, and the last 50 messages between two users are stored in Redis.
The subsequent message is sent to the recipient if they are also connected to the server.

## DIAGRAM
![Alt text](./doc/diagram.png?raw=true "High level diagram")
