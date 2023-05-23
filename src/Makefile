SERVER = build/candy-server
CLIENT = build/candy-client
.PHONY: server client
build: clean client server

server:
	go build -o $(SERVER) cmd/candy-server-server/main.go

client:
	go build -o $(CLIENT) cmd/candy-client/main.go

server_run:
	./build/candy-server --host localhost --tls-port 8080 --tls-certificate ./localhost/cert.pem --tls-key ./localhost/key.pem --scheme https

client_run:
	./build/candy-client -k AA -c 1 -m 20

clean:
	rm -rf build || true
