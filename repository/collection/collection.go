package collection

import (
	"context"
	"reflect"

	"github.com/architagr/repository/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ICollection[T any] interface {
	Disconnect()
	AddSingle(data T) (id interface{}, err error)
	AddMany(data []T) (ids []interface{}, err error)
	GetById(id interface{}) (data T, err error)
	Get(filter interface{}, pageSize int64, startPage int64) (data []T, err error)
}

type Collection[T any] struct {
	ctx         context.Context
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

func InitCollection[T any](conn connection.IConnection, databaseName, collection string) (ICollection[T], error) {
	client, ctx, err := conn.GetConnction()
	if err != nil {
		return nil, err
	}

	return &Collection[T]{
		ctx:         ctx,
		mongoClient: client,
		collection:  client.Database(databaseName).Collection(collection),
	}, nil
}

func (doc *Collection[T]) Disconnect() {
	doc.mongoClient.Disconnect(doc.ctx)
}

func (doc *Collection[T]) AddSingle(data T) (id interface{}, err error) {

	result, err := doc.collection.InsertOne(context.TODO(), data)
	if err != nil {
		return
	}
	id = result.InsertedID
	return
}

func (doc *Collection[T]) AddMany(data []T) (ids []interface{}, err error) {
	docs := make([]interface{}, len(data))
	for i, d := range data {
		docs[i] = d
	}
	result, err := doc.collection.InsertMany(context.TODO(), docs)
	if err != nil {
		return
	}
	ids = result.InsertedIDs
	return
}

func (doc *Collection[T]) GetById(id interface{}) (data T, err error) {
	filter := make(map[string]interface{})
	filter["_id"] = id
	result := doc.collection.FindOne(context.TODO(), filter)
	if result.Err() != nil {
		err = result.Err()
		return
	}
	resultDoc := new(T)
	err = result.Decode(resultDoc)
	if err != nil {
		return
	}
	data = *resultDoc
	return
}

func (doc *Collection[T]) Get(filter interface{}, pageSize int64, startPage int64) (data []T, err error) {
	if pageSize == 0 {
		pageSize = 10
	}
	skip := startPage * pageSize
	if skip > 0 {
		skip--
	}
	if filter == nil {
		filter = bson.D{{}}
	}
	filterOptions := options.Find()
	filterOptions.Limit = &pageSize
	filterOptions.Skip = &skip
	result, err := doc.collection.Find(context.TODO(), filter, filterOptions)
	if err != nil {
		return
	}
	data = make([]T, 0)
	err = result.All(context.TODO(), &data)
	if err != nil {
		return
	}
	return
}

func (doc *Collection[T]) interfaceSlice(slice interface{}) []T {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]T, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface().(T)
	}

	return ret
}
