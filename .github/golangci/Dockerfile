FROM golang:1.10.4

LABEL "com.github.actions.name"="build-test"
LABEL "com.github.actions.description"="run go test and build command"
LABEL "com.github.actions.icon"="terminal"
LABEL "com.github.actions.color"="blue"

LABEL "repository"="http://github.com/yogesh-desai/Hello_World"
LABEL "homepage"="http://github.com/yogesh-desai"
ENV GOROOT=/go
COPY entrypoint.sh /entrypoint.sh
RUN curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin v1.16.0

ENTRYPOINT ["/entrypoint.sh"]