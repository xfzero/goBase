package main

//使用mongo-driver包
//go get go.mongodb.org/mongo-driver/mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student1 struct {
	Name   string `bson: "name"`
	Age    int    `bson: "age"`
	Sid    string `bson: "sid"`
	Status int    `bson: "status"`
}

func main() {
	var (
		client     *mongo.Client
		err        error
		db         *mongo.Database
		collection *mongo.Collection
	)
	//1.建立连接
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Print(err)
		return
	}
	//2.选择数据库 my_db
	db = client.Database("test")

	//3.选择表 my_collection
	collection = db.Collection("student")

	//创建数据
	data := Student1{
		Name:   "seeta",
		Age:    18,
		Sid:    "s20180907",
		Status: 1,
	}

	//插入某一条数据
	if _, err = collection.InsertOne(context.TODO(), data); err != nil {
		fmt.Print(err)
	}

}
