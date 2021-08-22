FROM golang:1.16
WORKDIR /code
COPY . .
RUN make build
CMD ["event-server"]