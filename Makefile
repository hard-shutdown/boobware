build:
	go install mvdan.cc/garble@latest
	garble build -ldflags="-s -w" -trimpath	