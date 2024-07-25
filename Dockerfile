FROM ubuntu:latest
WORKDIR /app
COPY misgo ./
CMD ["./misgo"]