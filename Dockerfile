FROM arm32v7/golang:latest 
#RUN go get -v github.com/gin-gonic/gin
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

RUN env GOOS=linux GOARCH=arm GOARM=7 go build -o main . 
EXPOSE 7718
CMD ["/app/main"]

