build:
	CGO_ENABLED=1 go build -race -ldflags="-s -w" ./cmd/gbf 