package main

import (
	"log"

	jobcrawlerlinks "common-models/job-crawler-links"
	searchcondition "common-models/search-condition"

	"common-constants/constants"
	"repository/collection"
	"repository/connection"
)

type Jobs struct {
	JobId    int
	Title    string
	Company  string
	Location string
}

func main() {
	conn := connection.InitConnection("mongodb+srv://webscrapper:WebScrapper123@cluster0.xzvihm7.mongodb.net/?retryWrites=true&w=majority", 10)
	err := conn.ValidateConnection()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connection validated")

	doc, err := collection.InitCollection[jobcrawlerlinks.JobCrawlerLink](conn, "webscrapper", "jobLinks")
	if err != nil {
		log.Fatal(err)
	}

	// used to add links for orchestration
	links := []string{
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=25",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=50",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=75",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=100",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=150",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=175",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=200",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=225",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=250",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=275",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=300",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=325",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=350",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=375",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=400",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=425",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=450",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=475",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=500",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=525",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=550",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=575",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=600",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=625",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=650",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=675",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=700",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=725",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=750",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=775",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=800",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=725",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=850",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=875",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=900",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=925",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=950",
		"https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Digital%2BMarketing&location=India&locationId=&geoId=102713980&f_TPR=&f_E=2&f_WT=1&start=975",
	}

	for _, link := range links {
		doc.AddSingle(jobcrawlerlinks.JobCrawlerLink{
			HostName: constants.HostName_Linkedin,
			SearchCondition: searchcondition.SearchCondition{
				JobTitle:     "Digital Marketing",
				LocationInfo: searchcondition.Location{Country: "India", City: "any"},
				JobType:      "Full Time",
				JobModel:     "On site",
				RoleName:     "Digital Marketing",
				Experience:   "Entry Level"},
			Link: link,
		})
	}
	// id, err := doc.AddSingle(Jobs{JobId: 4, Title: "Workplace Services Manager, Google", Company: "Google", Location: "New York, NY"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("new id %+v", id)

	// newDoc, err := doc.GetById(id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("new doc retrived %+v", newDoc)

	// filter := map[string]interface{}{"title": "Workplace Services Manager, Google"}

	// result, err := doc.Get(filter, 10, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("filtered doc retrived %+v", result)

}
