FROM alpine:3.11.5
ENV LANG=en_US.UTF-8 \
    LANGUAGE=en_US.UTF-8



ENV TZ=Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone





WORKDIR /gowebframe
COPY out/* ./app
COPY web/ ./web



EXPOSE 8060
CMD ["/gowebframe/app"]