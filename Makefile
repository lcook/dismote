PREFIX?=	/usr/local

GO_CMD=		${PREFIX}/bin/go
GO_BIN?=	dismote
GO_FLAGS?=	-ldflags="-s -w"

default:
	${GO_CMD} build ${GO_FLAGS} -o ${GO_BIN}

build: default

clean:
	${GO_CMD} clean

mod:
	${GO_CMD} mod tidy -v
	${GO_CMD} mod verify

mod-update:
	${GO_CMD} get -u -v

update: mod-update mod

lint:
	${PREFIX}/bin/golangci-lint run

fmt:
	find . -name "*.go" -exec ${PREFIX}/bin/gofmt -w {} \;

.PHONY: build clean mod mod-update update lint fmt
