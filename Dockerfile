FROM golang:latest
WORKDIR /app/src/go_teste
ENV GOPATH=/app
COPY . /app/src/go_teste
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/handlers
RUN go build -o main .
CMD [ "./main" ]