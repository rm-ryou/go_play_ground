FROM golang:1.25 AS builder

WORKDIR /src

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o app ./main.go

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /

COPY --from=builder /src/app /app

ARG PORT=8080
ENV PORT=${PORT}

EXPOSE ${PORT}

USER nonroot:nonroot

CMD ["/app"]
