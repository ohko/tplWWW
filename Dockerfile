# 编译：docker build --no-cache -t tpler .
# 运行：docker run --rm -it -p 8080:8080 tpler
# 导出：docker save tpler > tpler.tar
# 导入：docker load < tpler.tar

FROM golang:1.15-buster AS builder
ENV GO111MODULE on
ENV CGO_ENABLED 1
ENV GOFLAGS -mod=vendor
ENV TZ Asia/Shanghai
COPY . /go/src
WORKDIR /go/src
RUN go build -v -o tpler_linux64 -ldflags "-s -w -X main.buildInfo='`git tag --contains master`.`git rev-parse --short HEAD`@`date '+%Y-%m-%d_%H:%M:%S_%Z_%z'`'" .

# ===========================

FROM debian:buster
LABEL maintainer="ohko <ohko@qq.com>"
COPY --from=builder /go/src/tpler_linux64 /
COPY dist/ /dist/
COPY public/ /public/
COPY view/ /view/
WORKDIR /
ENV TZ Asia/Shanghai
ENV LOG_LEVEL 1
EXPOSE 8080
ENTRYPOINT [ "/tpler_linux64" ]