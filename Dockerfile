FROM golang:latest
RUN apt-get update
RUN apt-get install -y bluez bluetooth

WORKDIR /go/src/github.com/paypal
RUN git clone https://github.com/paypal/gatt.git
WORKDIR /go/src/
RUN git clone https://github.com/edge7/linux_ble_server.git ble_rasbpi
WORKDIR ble_rasbpi
RUN go build -o ble_server main.go
CMD ["./ble_server"]