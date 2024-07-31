FROM ubuntu:latest
WORKDIR /app
COPY misgo ./
COPY templates ./templates
COPY static ./static
CMD ["./misgo"]