FROM golang:1.16.3-alpine

# install nodejs
RUN apk update \
    && apk add --no-cache \
    git \
    openssh-client

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
    && go install -tags 'dev' golang.org/x/tools/gopls@latest \
    && go install -tags 'dev' github.com/ramya-rao-a/go-outline@latest \
    && go install -tags 'dev' github.com/golang/mock/mockgen@latest \
    && go install -tags 'dev' golang.org/x/tools/cmd/godoc@latest
