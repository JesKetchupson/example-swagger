FROM golang:1.13 as app
WORKDIR /swagger-example
COPY . .
RUN  cd gen && go install cmd/example-swagger/main.go
EXPOSE 8080

CMD ["main"]