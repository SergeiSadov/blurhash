PROGRAM=blurhash
ENCODER=encoder
DECODER=decoder

all: help
help:
	@echo
	@echo "----------------------------------------------"
	@echo "encoder"
	@echo "build encoder"
	@echo
	@echo "decoder"
	@echo "build decoder encoder"
	@echo
	@echo "clean"
	@echo "remove binary files"
	@echo
	@echo "test"
	@echo "run tests"
	@echo "-----------------------------------------------"
	@echo

$(ENCODER):
		go build -o ${PROGRAM}-${ENCODER} cmd/${ENCODER}/main.go
$(ENCODER)_mac:
		GOOS=darwin GOARCH=amd64 go build -o ${PROGRAM}-${ENCODER} cmd/${ENCODER}/main.go

$(DECODER):
		go build -o ${PROGRAM}-${DECODER} cmd/${DECODER}/main.go
$(DECODER)_mac:
		GOOS=darwin GOARCH=amd64 go build -o ${PROGRAM}-${DECODER} cmd/${DECODER}/main.go

test:
		go test -v -coverprofile cover.out ./...

.PHONY: clean
clean:
		rm -f ${PROGRAM}-*