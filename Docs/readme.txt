
1) Using mongoDB (github.com/mongodb/mongo-go-driver/mongo)
Site: https://github.com/mongodb/mongo-go-driver

Install "dep" tool (golang/dep)
Site:
https://golang.github.io/dep/docs/introduction.html
https://github.com/golang/dep/blob/master/docs/Gopkg.toml.md

It might be required in package dir:
dep init

To install mongoDB driver:
dep ensure -add github.com/mongodb/mongo-go-driver/mongo