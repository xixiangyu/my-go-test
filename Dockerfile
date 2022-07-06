FROM golang as build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go env -w GOPROXY='https://goproxy.cn,direct' && go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./src/k8s_http

###
FROM golang as deploy
COPY --from=build /usr/local/bin/app /usr/local/bin/app
CMD ["app"]
