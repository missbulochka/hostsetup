FROM hostsetup-base:1.0 AS builder
WORKDIR /workspace/hostsetup
COPY . .
RUN go build -o /go/bin/res ./cmd/main.go

FROM ubuntu:22.04 AS runner
WORKDIR /app
COPY --from=builder /go/bin/res .
ENTRYPOINT [ "./res" ]