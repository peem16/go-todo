package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepositoryDB struct {
	*mongo.Collection
}

func NewRepositoryDB(db *mongo.Database) TodoRepository {
	collection := db.Collection("todos")

	return &TodoRepositoryDB{Collection: collection}
}

func (s *TodoRepositoryDB) Create(t Todo) error {
	_, err := s.Collection.InsertOne(context.Background(), t)
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
