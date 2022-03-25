package databases

import (
	"context"
	"log"
	"sync"

	"github.com/peam1146/todo_api/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Databases interface {
	InsertTodo(data bson.M) (primitive.ObjectID, error)
	GetAllTodos(result interface{}) error
	UpdateTodo(id primitive.ObjectID, data bson.M) error
	DeleteTodo(id primitive.ObjectID) error
	Close()
}

type databases struct {
	client *mongo.Client
}

var db Databases
var lock = &sync.Once{}

func InitDB() {
	log.Println("Initializing MongoDB client")
	uri := utils.Getenv("MONGODB_URI", "")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}

	usr := utils.Getenv("MONGODB_USER", "")
	pwd := utils.Getenv("MONGODB_PASSWORD", "")
	cred := options.Credential{}
	if usr != "" && pwd != "" {
		cred.Username = usr
		cred.Password = pwd
	}

	initClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetAuth(cred))
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary
	if err := initClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected and pinged.")
	db = &databases{
		client: initClient,
	}
}

func GetClient() Databases {
	lock.Do(func() {
		InitDB()
	})
	return db
}

func (d *databases) InsertTodo(data bson.M) (primitive.ObjectID, error) {
	collection := d.client.Database("todo").Collection("todos")
	res, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (d *databases) GetAllTodos(result interface{}) error {
	collection := d.client.Database("todo").Collection("todos")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	defer cur.Close(context.TODO())

	cur.All(context.TODO(), result)
	return nil
}

func (d *databases) UpdateTodo(id primitive.ObjectID, data bson.M) error {
	collection := d.client.Database("todo").Collection("todos")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": data})

	return err
}

func (d *databases) DeleteTodo(id primitive.ObjectID) error {
	collection := d.client.Database("todo").Collection("todos")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	return err
}

func (d *databases) Close() {
	if err := d.client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
