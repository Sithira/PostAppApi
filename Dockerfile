FROM golang:latest as build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /post-app-backend cmd/main.go

CMD /post-app-backend

FROM build as test

FROM build as prod
COPY --from=build /post-app-backend /app/post-app-backend
WORKDIR /app

RUN chmod +x post-app-backend

ENTRYPOINT ["/post-app-backend"]

EXPOSE 9000
