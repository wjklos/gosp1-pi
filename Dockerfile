FROM arm32v7/golang:latest 
RUN go get -v github.com/gin-gonic/gin
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
EXPOSE 7718 
RUN GOOS=linux GOARCH=arm7 go build -o main . 

CMD ["/app/main"]
