module orchestration-lambda

go 1.20

require (
	common-constants v0.0.0-00010101000000-000000000000
	common-models v0.0.0-00010101000000-000000000000
	mongodbRepo v0.0.0-00010101000000-000000000000
	github.com/aws/aws-lambda-go v1.39.1
	github.com/aws/aws-sdk-go v1.44.229
	github.com/aws/jsii-runtime-go v1.80.0
)

require (
	github.com/Masterminds/semver/v3 v3.2.0 // indirect
	github.com/PuerkitoBio/goquery v1.8.1 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/antchfx/htmlquery v1.3.0 // indirect
	github.com/antchfx/xmlquery v1.3.15 // indirect
	github.com/antchfx/xpath v1.2.3 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly/v2 v2.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/saintfish/chardet v0.0.0-20230101081208-5e3ef4b5456d // indirect
	github.com/temoto/robotstxt v1.1.2 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	github.com/yuin/goldmark v1.4.13 // indirect
	go.mongodb.org/mongo-driver v1.11.1 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
)

replace (
	common-constants => ./../common-constants
	common-models => ./../common-models
	mongodbRepo => ./../mongodbRepo
)
