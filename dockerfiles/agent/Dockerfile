FROM golang:1.19.4 AS builder
WORKDIR /
COPY ./integrator ./
RUN rm -rf vendor

RUN go mod download
RUN CGO_ENABLED=0 go build -o ./build/agent agent/cmd/agent/main.go

FROM scratch
COPY --from=builder ./build/agent agent

USER 65532:65532
ENTRYPOINT ["./agent"]
