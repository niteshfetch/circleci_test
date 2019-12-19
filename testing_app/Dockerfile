# build stage
FROM golang:alpine AS build-env
ADD src /src
RUN cd /src && go build -o testing-app main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/testing-app /app/
EXPOSE 8080
ENTRYPOINT ./testing-app