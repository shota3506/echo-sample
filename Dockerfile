FROM golang:latest

RUN go get github.com/labstack/echo
RUN go get github.com/jinzhu/gorm
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/go-sql-driver/mysql

WORKDIR /app
ADD . /app

CMD ["go", "run", "main.go"]