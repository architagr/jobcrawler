package persistance

import jobdetails "common-models/job-details"

type IPersistance interface {
	SaveMany(data []jobdetails.JobDetails) (ids []interface{}, err error)
	SaveSingle(data jobdetails.JobDetails) (id interface{}, err error)
}
