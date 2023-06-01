package persistance

import (
	jobdetails "common-models/job-details"
	"testing"
)

type mockCollection[T any] struct {
}

func (coll *mockCollection[T]) Disconnect() {}
func (coll *mockCollection[T]) AddSingle(data T) (id interface{}, err error) {
	id, err = 1, nil
	return
}
func (coll *mockCollection[T]) AddMany(data []T) (ids []interface{}, err error) {
	ids, err = []interface{}{1}, nil
	return
}
func (coll *mockCollection[T]) GetById(id interface{}) (data T, err error) {
	data, err = *new(T), nil
	return
}
func (coll *mockCollection[T]) Get(filter interface{}, pageSize int64, startPage int64) (data []T, err error) {
	data, err = []T{*new(T)}, nil
	return
}

type mockConfig struct {
}

func (e *mockConfig) GetDatabaseConnectionString() string {
	return "databaseConnectionString"
}
func (e *mockConfig) GetDatabaseName() string {
	return "databaseName"
}
func (e *mockConfig) GetCollectionName() string {
	return "collectionName"
}

func (e *mockConfig) IsLocal() bool {
	return true
}

func TestMongoPersistance(t *testing.T) {
	env := new(mockConfig)
	coll := new(mockCollection[jobdetails.JobDetails])
	t.Run("test persistance obj init", func(tb *testing.T) {
		obj := InitMongoDbPersistance(env, coll)
		if obj == nil {
			tb.Errorf("error in creating mongodb Persistance obj")
		}
	})

	t.Run("test persistance obj init", func(tb *testing.T) {
		obj := InitMongoDbPersistance(env, coll)
		_, err := obj.SaveMany([]jobdetails.JobDetails{
			{
				Title: "",
			},
		})
		if err != nil {
			tb.Errorf("error in savinf multiple obj in db")
		}
	})

	t.Run("test persistance obj init", func(tb *testing.T) {
		obj := InitMongoDbPersistance(env, coll)
		_, err := obj.SaveSingle(jobdetails.JobDetails{
			Title: "",
		})
		if err != nil {
			tb.Errorf("error in savinf multiple obj in db")
		}
	})
}
