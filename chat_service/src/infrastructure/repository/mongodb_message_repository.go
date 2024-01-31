package repository

import (
	"chat_service/src/domain/aggregate"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongodbMessageRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

type documentId struct {
	ID *primitive.ObjectID `bson:"_id"`
}

type document struct {
	documentId
	SenderId   int64     `bson:"sender_id"`
	ReceiverId int64     `bson:"receiver_id"`
	SentAt     time.Time `bson:"sent_at"`
	Content    string    `bson:"content"`
}

func NewMongodbMessageRepository(mongodb *mongo.Client) *MongodbMessageRepository {
	db := mongodb.Database("chat_db")

	return &MongodbMessageRepository{
		db:         db,
		collection: db.Collection("messages"),
	}
}

func (r *MongodbMessageRepository) HasConversationStarted(senderId int64, receiverID int64) bool {
	filter := bson.M{"sender_id": senderId, "receiver_id": receiverID}
	projection := bson.M{"_id": 1}

	var doc documentId
	r.collection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&doc)

	if doc.ID == nil {
		return false
	}

	return true
}

func (r *MongodbMessageRepository) AddMessage(message aggregate.Message) error {
	newMessage := bson.M{
		"sender_id":   message.SenderId,
		"receiver_id": message.ReceiverId,
		"content":     message.Message,
		"sent_at":     time.Now(),
	}

	_, err := r.collection.InsertOne(context.TODO(), newMessage)
	return err
}

func (r *MongodbMessageRepository) ListRecentMessages(senderId int64, receiverId int64, limit int64) []aggregate.Message {
	sort := bson.D{{"_id", -1}}
	opts := options.Find().SetSort(sort).SetLimit(limit)
	filter := bson.M{
		"sender_id":   senderId,
		"receiver_id": receiverId,
	}

	cursor, err := r.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var recentMessages []aggregate.Message
	for cursor.Next(context.TODO()) {
		var doc document
		err := cursor.Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}

		recentMessages = append(recentMessages, aggregate.Message{
			SenderId:   doc.SenderId,
			ReceiverId: doc.ReceiverId,
			SentAt:     doc.SentAt,
			Message:    doc.Content,
		})
	}

	return recentMessages
}
