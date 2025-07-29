#! /bin/bash

go mod init $1
go get  github.com/gin-gonic/gin
go get  github.com/go-playground/validator/v10
go get  github.com/golang-jwt/jwt/v5
go get	github.com/sirupsen/logrus
go get	github.com/spf13/viper
go get github.com/lib/pq
go get go.mongodb.org/mongo-driver/mongo