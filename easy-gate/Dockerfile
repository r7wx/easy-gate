FROM golang:1.18 AS go-builder
WORKDIR /opt/easy-gate
COPY . .
RUN make

FROM busybox:stable AS easy-gate
WORKDIR /usr/bin
COPY ./assets/easy-gate.json /etc/easy-gate/easy-gate.json
COPY --from=go-builder /opt/easy-gate/build/easy-gate .
ENTRYPOINT ["/usr/bin/easy-gate", "/etc/easy-gate/easy-gate.json"]