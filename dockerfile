# FROM golang:1.13-alpine

# ENV GO111MODULE=off
# ENV PROJECT_PATH=github.com/Frosin/shoplist-api-client-go

# ARG _path
# RUN apk add --no-cache --update \
# git
# RUN apk add libc-dev
# RUN apk add gcc
# ENV GOPATH=/go \
# PATH="/go/bin:$PATH"
# RUN mkdir -p ${GOPATH}/src/${PROJECT_PATH}
# WORKDIR ${GOPATH}/src/${PROJECT_PATH}
# COPY . .

# WORKDIR ${_path}
# RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=develop" -o bin/shoplist .

# EXPOSE 80