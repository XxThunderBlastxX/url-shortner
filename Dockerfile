# syntax=docker/dockerfile:1

FROM golang:1.16.6 AS api

WORKDIR /app

COPY go.mod go.sum ./
COPY .env ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortner ./main.go
#
## Use 'scratch' image for super-mini build.
#FROM scratch AS prod
#
## Set working directory for this stage.
#WORKDIR /production
#
## Copy our compiled executable from the last stage.
#COPY --from=api /compiler/url-shortner .

# Run application and expose port 8080.
EXPOSE 3200
CMD ["./url-shortner"]