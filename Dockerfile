FROM golang:1.19


COPY . .

RUN go build -o app app.go

ENTRYPOINT [ "./app" ]

