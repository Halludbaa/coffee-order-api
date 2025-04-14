#! /bin/bash

go mod init $1
go get  github.com/gin-gonic/gin v1.10.0
go get  github.com/go-playground/validator/v10 v10.26.0
go get  github.com/golang-jwt/jwt/v5 v5.2.2
go get	github.com/sirupsen/logrus v1.9.3
go get	github.com/spf13/viper v1.20.1