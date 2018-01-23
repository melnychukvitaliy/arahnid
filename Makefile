HOST=192.168.1.1

all: clean build run

build:
	go get gobot.io/x/gobot
	go get gobot.io/x/gobot/drivers/gpio
	go get gobot.io/x/gobot/platforms/raspi

	GOOS=linux GOARCH=arm GOARM=5 go build -o arahnid

run:
	#run over ssh

clean:
	rm ./arahnid

deploy:
	scp ./arahnid pi@$(HOST) /home/pi