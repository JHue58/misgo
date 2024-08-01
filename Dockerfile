FROM ubuntu:latest
WORKDIR /app
RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install -y tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY misgo ./
COPY templates ./templates
COPY static ./static
CMD ["./misgo"]