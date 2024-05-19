FROM golang:1.22 AS go-builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
GOPROXY=https://goproxy.cn,direct \
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64

WORKDIR /build

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum /build/
RUN go mod download && go mod verify

COPY . .
RUN go build -o srbbs-app ./src/

CMD ["srbbs-app"]

EXPOSE 8081

ENTRYPOINT ["srbbs-app"]
