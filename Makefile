build:
	go build -o udp-relay -v

clean:
	rm -f udp-relay

run:
	go run . -localHost 127.0.0.1 -localPort 65321 -remoteHost 127.0.0.1 -remotePort 65322
