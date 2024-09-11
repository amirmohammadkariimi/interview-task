FROM golang:1.23.1-alpine3.20 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
ADD . .
RUN go build -o interview_task ./cmd/

FROM alpine:3.19
COPY --from=builder /app/interview_task /bin/interview_task
EXPOSE 3000
ENTRYPOINT ["/bin/interview_task"]