From golang:1.10.1-alpine
ENV SOURCES /go/src/
# ENV GOPATH /usr/local/go/
ENV GOBIN /go/bin
COPY ./src ${SOURCES}
# # RUN chmod 777 -R ${SOURCES}
RUN go install src/microservice.go
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT /go/bin/microservice
