FROM golang
COPY . .
RUN go build -o main .
EXPOSE 9800
CMD ["./main"]
