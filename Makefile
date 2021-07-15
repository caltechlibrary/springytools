
PROJECT = springytools

VERSION = $(shell jq .version codemeta.json | sed -E 's/"//g')

PROGRAMS = lgxml2json linkreport

PACKAGE = $(shell ls -1 *.go)

#FIXME: Need to make PREFIX conditional on OS, especially Windows
#PREFIX = /usr/local/bin
PREFIX = $(HOME)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
        EXT = .exe
endif

ifneq ($(prefix),)
	PREFIX = $(prefix)
endif

build: version.go $(PROGRAMS)

version.go: .FORCE
	@echo "package $(PROJECT)" >version.go
	@echo '' >>version.go
	@echo 'const Version = "$(VERSION)"' >>version.go
	@echo '' >>version.go

$(PROGRAMS): $(PACKAGE)
	mkdir -p bin
	go build -o bin/$@$(EXT) cmd/$@/$@.go

test: $(PACKAGE)
	go test

install: build $(PROGRAMS)
	@echo "Installing programs in $(PREFIX)/bin"
	@cp -vR bin/* $(PREFIX)/bin/
	@echo ""
	@echo "Make sure $(PREFIX)/bin is in your PATH"

clean: .FORCE
	@if [ -f version.go ]; then rm version.go; fi
	@if [ -d bin ]; then rm -fR bin; fi
	@if [ -d dist ]; then rm -fR dist; fi
	@if [ -d man ]; then rm -fR man; fi


dist-documentation: .FORCE
	cp LICENSE dist/
	cp codemeta.json dist/
	cp CITATION.cff dist/
	cp [A-Z]*.md dist/

dist-Linux-x86_64: .FORCE
	@if [ -d dist/bin ]; then rm -fR dist/bin; fi
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/lgxml2json cmd/lgxml2json/lgxml2json.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/linkreport cmd/linkreport/linkreport.go
	cd dist && zip -r $(PROJECT)-Linux-x86_64-$(VERSION).zip LICENSE *.json *.cff *.md bin/*

dist-Darwin-x86_64: .FORCE
	@if [ -d dist/bin ]; then rm -fR dist/bin; fi
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/lgxml2json cmd/lgxml2json/lgxml2json.go
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/linkreport cmd/linkreport/linkreport.go
	cd dist && zip -r $(PROJECT)-macOS-x86_64-$(VERSION).zip LICENSE *.json *.cff *.md bin/*

dist-Darwin-arm64: .FORCE
	@if [ -d dist/bin ]; then rm -fR dist/bin; fi
	env GOOS=darwin GOARCH=arm64 go build -o dist/bin/lgxml2json cmd/lgxml2json/lgxml2json.go
	env GOOS=darwin GOARCH=arm64 go build -o dist/bin/linkreport cmd/linkreport/linkreport.go
	cd dist && zip -r $(PROJECT)-macOS-arm64-$(VERSION).zip LICENSE *.json *.cff *.md bin/*

dist-Raspbian-arm7: .FORCE
	@if [ -d dist/bin ]; then rm -fR dist/bin; fi
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/lgxml2json cmd/lgxml2json/lgxml2json.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/linkreport cmd/linkreport/linkreport.go
	cd dist && zip -r $(PROJECT)-Raspbian-arm7-$(VERSION).zip LICENSE *.json *.cff *.md bin/*

dist-Windows-x86_64: .FORCE
	@if [ -d dist/bin ]; then rm -fR dist/bin; fi
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/lgxml2json.exe cmd/lgxml2json/lgxml2json.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/linkreport.exe cmd/linkreport/linkreport.go
	cd dist && zip -r $(PROJECT)-Windows-x86_64-$(VERSION).zip LICENSE *.json *.cff *.md bin/*


dist: .FORCE
	mkdir -p dist

release: dist version.go $(PROGRAMS) dist-documentation dist-Linux-x86_64 dist-Raspbian-arm7 dist-Darwin-arm64 dist-Darwin-x86_64 dist-Windows-x86_64

.FORCE:
