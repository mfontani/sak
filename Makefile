sak: *.go go.mod
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -tags timetzdata --ldflags "-X 'main.Version=$(shell git describe --tags 2>/dev/null || git rev-parse HEAD 2>/dev/null || echo unknown)' -extldflags \"-static\" -s -w" -o sak .
