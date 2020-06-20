FROM golang:1.14 as builder
EXPOSE 8080
COPY cmd /build/cmd
COPY go.mod go.sum /build/
RUN go env
# COPY go.mod go.sum /build/
WORKDIR /build
RUN useradd scratchuser && \
    export GOPATH="" && go mod vendor && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /tasker-backend ./cmd/

# FROM scratch
# COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
# COPY --from=builder /tasker-backend /tasker-backend
# COPY --from=builder /etc/passwd /etc/passwd
# USER scratchuser
# CMD /tasker-backend
ENTRYPOINT ["/tasker-backend"]