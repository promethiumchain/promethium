# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: promethium android ios promethium-cross evm all test clean
.PHONY: promethium-linux promethium-linux-386 promethium-linux-amd64 promethium-linux-mips64 promethium-linux-mips64le
.PHONY: promethium-linux-arm promethium-linux-arm-5 promethium-linux-arm-6 promethium-linux-arm-7 promethium-linux-arm64
.PHONY: promethium-darwin promethium-darwin-386 promethium-darwin-amd64
.PHONY: promethium-windows promethium-windows-386 promethium-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

promethium:
	build/env.sh go run build/ci.go install ./cmd/promethium
	@echo "Done building."
	@echo "Run \"$(GOBIN)/promethium\" to launch promethium."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/promethium.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/promethium.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

lint: ## Run linters.
	build/env.sh go run build/ci.go lint

clean:
	./build/clean_go_build_cache.sh
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

promethium-cross: promethium-linux promethium-darwin promethium-windows promethium-android promethium-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/promethium-*

promethium-linux: promethium-linux-386 promethium-linux-amd64 promethium-linux-arm promethium-linux-mips64 promethium-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/promethium-linux-*

promethium-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/promethium
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/promethium-linux-* | grep 386

promethium-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/promethium
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/promethium-linux-* | grep amd64

promethium-linux-arm: promethium-linux-arm-5 promethium-linux-arm-6 promethium-linux-arm-7 promethium-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/promethium-linux-* | grep arm

promethium-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/promethium
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/promethium-linux-* | grep arm-5

promethium-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/promethium
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/promethium-linux-* | grep arm-6

promethium-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/promethium
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/promethium-linux-* | grep arm-7

promethium-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/promethium
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/promethium-linux-* | grep arm64

promethium-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/promethium
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/promethium-linux-* | grep mips

promethium-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/promethium
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/promethium-linux-* | grep mipsle

promethium-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/promethium
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/promethium-linux-* | grep mips64

promethium-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/promethium
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/promethium-linux-* | grep mips64le

promethium-darwin: promethium-darwin-386 promethium-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/promethium-darwin-*

promethium-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/promethium
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/promethium-darwin-* | grep 386

promethium-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/promethium
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/promethium-darwin-* | grep amd64

promethium-windows: promethium-windows-386 promethium-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/promethium-windows-*

promethium-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/promethium
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/promethium-windows-* | grep 386

promethium-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/promethium
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/promethium-windows-* | grep amd64
