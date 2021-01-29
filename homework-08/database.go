package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	Mongo `inline`
	Autor string `bson:"title"`
	Date string `bson:"date"`
	Header string `bson:"header"`
	Content string `bson:"content"`
}

func (p *Post) GetMongoCollectionName() string {
	return "jokes"
}

func (p *Post) Insert(ctx context.Context, db *mongo.Database) error {
	p.ID = primitive.NewObjectID()
	coll := db.Collection(p.GetMongoCollectionName())
	_, err := coll.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

func GetPost(ctx context.Context, db *mongo.Database, id primitive.ObjectID) (*Post, error) {
	p := Post{}
	coll := db.Collection(p.GetMongoCollectionName())
	res := coll.FindOne(ctx, bson.M{"_id": id})
	if err := res.Decode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Post) Update(ctx context.Context, db *mongo.Database) error {
	coll := db.Collection(p.GetMongoCollectionName())
	_, err := coll.ReplaceOne(ctx, bson.M{"_id": p.ID}, p)
	return err
}

func (p *Post) Delete(ctx context.Context, db *mongo.Database) error {
	coll := db.Collection(p.GetMongoCollectionName())
	_, err := coll.DeleteOne(ctx, bson.M{"_id": p.ID})
	return err
}

func GetPosts(ctx context.Context, db *mongo.Database) ([]Post, error) {
	p := Post{}
	coll := db.Collection(p.GetMongoCollectionName())

	cur, err := coll.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	posts := []Post{}
	if err := cur.All(ctx, &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func Find(ctx context.Context, db *mongo.Database, field string, value interface{}) ([]Post, error) {
	p := Post{}
	coll := db.Collection(p.GetMongoCollectionName())

	cur, err := coll.Find(ctx, bson.M{field: value})
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	if err := cur.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
