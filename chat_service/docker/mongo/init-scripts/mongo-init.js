db = db.getSiblingDB('chat_db');
db.createUser({
    user: 'admin',
    pwd: 'admin123',
    roles: [
        {
            role: 'readWrite',
            db: 'chat_db',
        },
    ],
});

db.createCollection('messages');

db.messages.createIndex(
    { sender_id: 1, receiver_id: 1 },
    { name: 'sender_receiver_index' }
);