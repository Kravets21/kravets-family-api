FROM golang:1.17-buster as builder

WORKDIR /go/src/app

COPY . .

RUN go build -v -o api cmd/api/main.go

FROM builder as prod

WORKDIR /bin

COPY --from=builder /go/src/app/api .

CMD api