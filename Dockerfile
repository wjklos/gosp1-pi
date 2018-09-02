FROM arm32v7/golang:latest 
RUN go get -v github.com/gin-gonic/gin
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
EXPOSE 7718 

ENV GOOS=linux
ENV GOARCH=arm
ENV GOARM=7

RUN go build -o main . 

CMD ["/app/main"]

