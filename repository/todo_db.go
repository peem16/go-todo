package repository

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TodoRepositoryDB struct {
	*mongo.Collection
}

func NewRepositoryDB(db *mongo.Database) TodoRepository {
	collection := db.Collection("todos")
	mod := mongo.IndexModel{
		Keys: bson.M{
			"todoid": 1, // index in ascending order
		}, Options: options.Index().SetUnique(true),
	}
	ind, err := collection.Indexes().CreateOne(context.Background(), mod)

	// Check if the CreateOne() method returned any errors
	if err != nil {
		fmt.Println("Indexes().CreateOne() ERROR:", err)
		os.Exit(1) // exit in case of error
	} else {
		// API call returns string of the index name
		fmt.Println("CreateOne() index:", ind)
	}

	return &TodoRepositoryDB{Collection: collection}
}

func (repo *TodoRepositoryDB) Count() (int64, error) {
	count, err := repo.Collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}

	return count, err
}

func (repo *TodoRepositoryDB) Create(t Todo) error {
	_, err := repo.Collection.InsertOne(context.Background(), t)
	return err
}

func (s *TodoRepositoryDB) GetAll() ([]Todo, error) {
	cur, err := s.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	todoList := []Todo{}
	todo := Todo{}

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {

		err := cur.Decode(&todo)
		if err != nil {
			return nil, err
		}

		todoList = append(todoList, todo)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return todoList, nil
}

func (s *TodoRepositoryDB) GetByID(id int) (*Todo, error) {
	todo := &Todo{}
	filter := bson.D{{"todoid", 1}}
	err := s.Collection.FindOne(context.Background(), filter).Decode(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoRepositoryDB) UpdateByID(id int, todo Todo) error {
	update := bson.M{
		"$set": bson.M{"title": todo.Title, "status": todo.Status},
	}
	filter := bson.D{{Key: "todoID", Value: id}}

	_, err := s.Collection.UpdateOne(context.Background(), filter, update, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoRepositoryDB) Delete(id int) error {

	filter := bson.D{{Key: "todoID", Value: id}}

	_, err := s.Collection.DeleteOne(context.Background(), filter, nil)
	if err != nil {
		return err
	}

	return nil
}
