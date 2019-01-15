all: clean build run

build:
	go get gobot.io/x/gobot
	go get gobot.io/x/gobot/drivers/gpio
	go get gobot.io/x/gobot/platforms/firmata

	go build -o arahnid

run:
	./arahnid

clean:
	rm -f ./arahnid

	