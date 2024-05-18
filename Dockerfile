#FROM ubuntu:latest
#LABEL authors="hz"
#
#ENTRYPOINT ["top", "-b"]

FROM golang:1.22 AS go-builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
GOPROXY=https://goproxy.io,direct \
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64

WORKDIR /build

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o srbbs-app

CMD ["srbbs-app"]


FROM postgres:14.3 AS postgres-builder
RUN localedef -i de_DE -c -f UTF-8 -A /usr/share/locale/locale.alias de_DE.UTF-8
ENV LANG de_DE.utf8
ENV POSTGRES_USER srbbs
ENV POSTGRES_PASSWORD 123456
COPY ./script/setup.sql /docker-entrypoint-initdb.d


#FROM redis AS redis-builder
#COPY redis.conf /usr/local/etc/redis/redis.conf
#CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]



# 基础镜像 从前两个构建阶段复制生成的可执行文件
FROM centos:7
ENV container docker
RUN (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == \
systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;
VOLUME [ "/sys/fs/cgroup" ]
CMD ["/usr/sbin/init"]



# 从builder镜像中把可执行文件拷贝到当前目录
COPY --from=go-builder /build  /


# 声明服务端口
EXPOSE 8081

# 需要运行的命令
ENTRYPOINT ["/srbbs-app"]
