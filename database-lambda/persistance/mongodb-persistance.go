package persistance

import (
	jobdetails "common-models/job-details"
	"database_lambda/config"
	"mongodbRepo/collection"
)

type MongodbPersistance struct {
	doc collection.ICollection[jobdetails.JobDetails]
}

func InitMongoDbPersistance(env config.IConfig, doc collection.ICollection[jobdetails.JobDetails]) IPersistance {
	return &MongodbPersistance{
		doc: doc,
	}
}
func (persistance *MongodbPersistance) SaveMany(data []jobdetails.JobDetails) (ids []interface{}, err error) {
	ids, err = persistance.doc.AddMany(data)
	return
}

func (persistance *MongodbPersistance) SaveSingle(data jobdetails.JobDetails) (id interface{}, err error) {
	id, err = persistance.doc.AddSingle(data)
	return
}
