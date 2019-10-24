PREFIX = /usr
GOSITE_DIR = ${PREFIX}/share/gocode
GOPKG_PERFIX = github.com/linuxdeepin/go-dbus-factory
SRC_DIR=${DESTDIR}${GOSITE_DIR}/src/${GOPKG_PERFIX}

all: build

build:
	echo ignore build

bin:
	go build -o generator _tool/generator/*.go

install:
	mkdir -p ${SRC_DIR}
	for dir in object_manager net.* org.* com.*;do\
		mkdir -p ${SRC_DIR}/$$dir;\
		cp $$dir/*.go ${SRC_DIR}/$$dir;\
	done

vet:
	go vet ./...
	go vet ./_tool/generator
	go vet ./_tool/prop_gen
