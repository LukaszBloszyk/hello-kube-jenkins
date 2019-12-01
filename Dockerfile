#comment more
FROM golang:1.13.4-alpine3.10
LABEL maintainer = "Lukasz Bloszyk lukasz.bloszyk@gazeta.pl"
RUN mkdir /go/src/hello-go-demo/
WORKDIR /go/src/hello-kubernetes-jenkins/
COPY . .
RUN go install -v
EXPOSE 8080
CMD ["hello-kubernetes-jenkins"]