
FROM golang:1.24.4-bullseye AS builder

WORKDIR /workdir

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_CPPFLAGS="-D_FORTIFY_SOURCE=2 -fstack-protector-all"
ENV GOFLAGS="-buildmode=pie"

RUN go build ./cmd/app


FROM gcr.io/distroless/base-debian11:nonroot

COPY --from=builder /workdir/app /app/app

COPY --from=builder /workdir/config.yaml /app/config.yaml

USER 65534

ENTRYPOINT ["/app/app"]
