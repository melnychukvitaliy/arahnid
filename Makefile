HOST=192.168.7.105

all: clean build deploy run

build:
	go get gobot.io/x/gobot
	go get gobot.io/x/gobot/drivers/gpio
	go get gobot.io/x/gobot/platforms/raspi

	GOOS=linux GOARCH=arm GOARM=5 go build -o arahnid

run:
	ssh -t pi@$(HOST) "./arahnid/arahnid"

clean:
	rm -f ./arahnid

deploy:
	scp ./arahnid pi@$(HOST):/home/pi/arahnid