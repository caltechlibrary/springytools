#
# Default Go project makefile, Caltech Library
#
PROJECT = springytools

VERSION = $(shell jq .version codemeta.json | sed -E 's/"//g')

PROGRAMS = $(shell ls -1 cmd/)

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
	@mkdir -p bin
	go build -o bin/$@$(EXT) cmd/$@/$@.go

test: $(PACKAGE)
	go test

install: build
	@echo "Installing programs in $(PREFIX)/bin"
	@for FNAME in $(PROGRAMS); do if [ -f ./bin/$$FNAME ]; then cp -v ./bin/$$FNAME $(PREFIX)/bin/$$FNAME; fi; done
	@echo ""
	@echo "Make sure $(PREFIX)/bin is in your PATH"

uninstall: .FORCE
	@echo "Removing programs in $(PREFIX)/bin"
	@for FNAME in $(PROGRAMS); do if [ -f $(PREFIX)/bin/$$FNAME ]; then rm -v $(PREFIX)/bin/$$FNAME; fi; done

status:
	@git status

save:
	@if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	@git push origin $(BRANCH)

refresh:
	@git fetch origin
	@git pull origin $(BRANCH)

clean: .FORCE
	@if [ -f version.go ]; then rm version.go; fi
	@if [ -d bin ]; then rm -fR bin; fi
	@if [ -d dist ]; then rm -fR dist; fi
	@if [ -d man ]; then rm -fR man; fi


dist-documentation: .FORCE
	@cp LICENSE dist/
	@cp codemeta.json dist/
	@cp CITATION.cff dist/
	@cp [A-Z]*.md dist/

dist-Linux-x86_64: .FORCE
	@if [ -d dist/bin ]; then rm -fR dist/bin; fi
	@for FNAME in $(PROGRAMS); do env GOOS=linux GOARCH=amd64 go build -o dist/bin/$$FNAME cmd/$$FNAME/$$FNAME.go; done
	@cd dist && zip -r $(PROJECT)-Linux-x86_64-$(VERSION).zip LICENSE *.json *.cff *.md bin/*

dist-Darwin-x86_64: .FORCE
	@if [ -d dist/bin ]; then rm -fR dist/bin; fi
	@for FNAME in $(PROGRAMS); do env GOOS=darwin GOARCH=amd64 go build -o dist/bin/$$FNAME cmd/$$FNAME/$$FNAME.go; done
	@cd dist && zip -r $(PROJECT)-macOS-x86_64-$(VERSION).zip LICENSE *.json *.cff *.md bin/*

dist-Darwin-arm64: .FORCE
	@if [ -d dist/bin ]; then rm -fR dist/bin; fi
	@for FNAME in $(PROGRAMS); do env GOOS=darwin GOARCH=arm64 go build -o dist/bin/$$FNAME cmd/$$FNAME/$$FNAME.go; done
	@cd dist && zip -r $(PROJECT)-macOS-arm64-$(VERSION).zip LICENSE *.json *.cff *.md bin/*

dist-Raspbian-arm7: .FORCE
	@if [ -d dist/bin ]; then rm -fR dist/bin; fi
	@for FNAME in $(PROGRAMS); do env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/$$FNAME cmd/$$FNAME/$$FNAME.go; done
	@cd dist && zip -r $(PROJECT)-Raspbian-arm7-$(VERSION).zip LICENSE *.json *.cff *.md bin/*

dist-Windows-x86_64: .FORCE
	@if [ -d dist/bin ]; then rm -fR dist/bin; fi
	@for FNAME in $(PROGRAMS); do env GOOS=windows GOARCH=amd64 go build -o dist/bin/$$FNAME.exe cmd/$$FNAME/$$FNAME.go; done
	@cd dist && zip -r $(PROJECT)-Windows-x86_64-$(VERSION).zip LICENSE *.json *.cff *.md bin/*


dist: .FORCE
	@mkdir -p dist

release: version.go dist dist-documentation dist-Linux-x86_64 dist-Raspbian-arm7 dist-Darwin-arm64 dist-Darwin-x86_64 dist-Windows-x86_64

.FORCE:
