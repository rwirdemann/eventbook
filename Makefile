build-linux:
	env GOOS=linux GOARCH=amd64 go build -o bin/eventbook main.go

deploy: build-linux
	ssh root@95.217.180.178 "pkill eventbook"
	scp bin/eventbook root@95.217.180.178:~
	ssh root@95.217.180.178 "sh -c 'nohup /root/eventbook > /dev/null 2>&1 &'"