FROM golang:1.18
#RUN hciconfig
#RUN hciconfig hci0 down  # or whatever hci device you want to use
#RUN service bluetooth stop
RUN apt-get update
RUN apt-get install -y bluez bluetooth
WORKDIR /home/edge7/
COPY . .
RUN echo $(ls)
RUN go build -o ble_server main.go
CMD ["./ble_server"]