FROM golang:latest

RUN go get github.com/labstack/echo
RUN go get github.com/jinzhu/gorm
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/go-sql-driver/mysql
RUN go get gopkg.in/go-playground/validator.v9

WORKDIR /go/src/app
ADD . /go/src/app

CMD ["go", "run", "main.go"]