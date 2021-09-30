/*

bson.D:
用来表示一个有序的 BSON 文档，当我们需要传递一些有序的参数的时候可以使用，比如 MongoDB 中的聚合操作，
聚合操作往往是一组操作的集合，比如先筛选、然后分组、然后求和，这个顺序是不能乱的。

bson.E:
用来表示 bson.D 中的一个属性，类型定义如下:
// E represents a BSON element for a D. It is usually used inside a D.
type E struct {
	Key   string
	Value interface{}
}
可以类比 json 对象里面的某一个属性以及其值。

bson.M:
用来表示无需的 BSON 文档，就是一个普通的 map 类型。在保存到 MongoDB 中的时候，字段顺序是不确定的，而 bson.D 的顺序是确定的。

顺序可能大部分情况都是不需要的。不过在匹配嵌套的 BSON 数组文档的时候，可能会有问题。但还是有其他的解决办法的。

bson.A:
用来表示 BSON 文档中的数组类型，是有序的。
*/

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	address   = "mongodb://localhost:27017"
	timeout   = 5 * time.Second
	mgoClient *mongo.Client
	dbName    = "test1"
)

type Student struct {
	Id   uint64
	Name string
	age  uint64
}

func connectMgo() {
	var err error
	clinetOptions := options.Client().ApplyURI(address).SetConnectTimeout(timeout)

	//链接到MongoDB
	mgoClient, err = mongo.Connect(context.TODO(), clinetOptions)
	if err != nil {
		log.Fatal(err)
	}

	//检查链接
	err = mgoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func closeMgo() bool {
	if mgoClient == nil {
		return true
	}

	err := mgoClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Connection to MongoDB closed.")

	return true
}

func getCollection(cName string) *mongo.Collection {
	return mgoClient.Database(dbName).Collection(cName)
}

func FindOne(cName string, filter interface{}, ret interface{}, opts ...*options.FindOneOptions) (not bool, err error) {
	err = getCollection(cName).FindOne(context.TODO(), filter).Decode(ret)
	if err != nil && err == mongo.ErrNoDocuments {
		not = true
		err = nil
	}

	return
}

func test1() {
	user := &Student{}

	not, _ := FindOne("students", bson.M{"name": "tony"}, user)
	if not {
		fmt.Println("查询失败")
		return
	}

	fmt.Println("user:", user)
}

func main() {
	connectMgo()
	defer closeMgo()

	test1()
}
