PREFIX = /usr
GOSITE_DIR = ${PREFIX}/share/gocode
GOPKG_PERFIX = github.com/linuxdeepin/go-dbus-factory
SRC_DIR=${DESTDIR}${GOSITE_DIR}/src/${GOPKG_PERFIX}

all: build

build:
	echo ignore build

print_gopath:
	GOPATH="${CURDIR}/${GOPATH_DIR}:${GOPATH}"

bin:
	cd _tool/generator && go build -o generator
	cp _tool/generator/generator .

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
