# Job Crawler

This project is intended to crawl jobs from different potals and have them in database.
for now we have pushed jobs to Mongodb, but can be pushed to any DB of choice

<hr/>
We have 6 major components :

## Orchestration

This component gets crawlable job listing links from the DB and pushes all links to the SNS that will start crawling.

## Crawler

We have different SQS listning to SNS where Orchestration pushes the links, each SQS is dedicated to host which we will crawl.
this is done to have parallel execution of jobs from different portals with the frequency there individual gateways allows.

Crawler crawls the listing page and gets the link to specific job and pushes to scrapper SNS.

## Scrapper

We have different SQS listning to SNS where Crawler pushes the specific job links, each SQS is dedicated to host which we will scrap.
This is again done to have parallel execution for jobs scrapping from different portals with the frequency there individual gateways allows.

This component scraps the job details from the link and pushes to the database SNS.

## Database

We have different SQS listning to SNS where Scrapper pushes the specific job details, each SQS is dedicated to host which we will scrap.

This component pushes data to the db, this component was taken out of scrapper so that we can hve better control on the data model, and how we will push data in the DB.
We push data to a temp table in this component. so that switch table component can backup exiting jobs and make this temp table as main table to be used by clients.

## Monitoring

This component keep a track on all the SQS to see if the scrapping, crawling and saving to db is completed. if so we then we init the switch table component.

## Switch Table

This component creates backup of existing jobs and replaces the newly fetched jobs to main table to be used by clients.
