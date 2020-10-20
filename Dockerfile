FROM golang:latest
hciconfig
hciconfig hci0 down  # or whatever hci device you want to use
service bluetooth stop

WORKDIR /go/src/github.com/paypal
RUN git clone https://github.com/paypal/gatt.git
WORKDIR /go/src/
RUN git clone https://github.com/edge7/linux_ble_server.git
WORKDIR linux_ble_server
RUN go build -o ble_server main.go
./ble_server
