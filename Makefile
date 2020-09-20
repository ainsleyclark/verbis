build:
	go build  ./main.go

run:
	go run ./main.go

serve:
	go run ./main.go serve

install-serve:
	go install cms && cms serve

live-serve:
	HOST="localhost" gin -i --port=8080 --laddr=127.0.0.1 --all run serve

live-test:
	HOST="localhost" gin -i --port=8080 --laddr=127.0.0.1 --all run test

install:
	go install cms