APP-NAME = sshtgbot

download:
	go mod download

build:
	go build .

setup:
	sudo mkdir -p /usr/local/bin/$(APP-NAME)/
	sudo cp ssh_tg_bot /usr/local/bin/$(APP-NAME)/
	sudo cp config.json /usr/local/bin/$(APP-NAME)/
	sudo cp ssh_notification.sh /etc/ssh/sshrc
	sudo chown root:root /etc/ssh/sshrc
	sudo chmod 0755 /etc/ssh/sshrc

run:
	go run .

test:
	go test