#!/bin/sh

docker pull registry.cn-chengdu.aliyuncs.com/zx-tech/tpler
docker rm -fv tpler && \
docker run -d --name=tpler --restart=always \
   -p 8079:8080 \
   -v /data/ebayerp/data/session:/session \
   -v /usr/share/zoneinfo:/usr/share/zoneinfo:ro \
   -v /etc/ssl/certs:/etc/ssl/certs:ro \
   -e TZ=Asia/Shanghai \
   -e LOG_LEVEL=1 \
   --link postgres:postgres \
   registry.cn-chengdu.aliyuncs.com/zx-tech/tpler