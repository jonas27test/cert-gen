FROM golang as builder
EXPOSE 8080
COPY . .
RUN useradd scratchuser && \
    export GOPATH="" && go mod vendor && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /cert-gen ./src/

FROM scratch
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
COPY --from=builder /cert-gen /cert-gen
COPY --from=builder /etc/passwd /etc/passwd
USER scratchuser
CMD ["/cert-gen"]