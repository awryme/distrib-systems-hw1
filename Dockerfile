FROM golang:1.13 as builder
COPY . /app
WORKDIR /app
ENV GOPROXY=https://gocenter.io
RUN go build -v -o inc
RUN ls
RUN pwd

from debian
COPY --from=builder /app/inc /app/inc
CMD [ "/app/inc" ]