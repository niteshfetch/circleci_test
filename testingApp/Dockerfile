# build stage
FROM golang:alpine AS build-env
ADD src /src
RUN cd /src && go build -o testingApp main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/testingApp /app/
EXPOSE 8080
ENTRYPOINT ./testingApp