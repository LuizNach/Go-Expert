# BuildStage

FROM golang:1.21.10-alpine AS build-stage

WORKDIR /app

COPY server.go .
COPY go.mod .
COPY go.sum .

RUN go mod tidy

# Reference: https://www.youtube.com/watch?v=cGYfMIKHC30&t=493s&ab_channel=GoLabTutoriais

RUN GOOS=linux go build -o my_app server.go

# Deploy Stage

FROM scratch

WORKDIR /app

COPY --from=build-stage /app/my_app ./

EXPOSE 8080

ENTRYPOINT ["./my_app"]