export CONF_PATH=./config/dev.env

rbd:
	go build -tags debug -o ./.build/auto_order.exe ./cmd
	./.build/auto_order.exe

rb:
	go build -o ./.build/auto_order.exe ./cmd
	./.build/auto_order.exe

build:
	go build -o ./.build/auto_order.exe ./cmd

run:
	./.build/auto_order.exe

test:
	echo "Hello Makefile!"