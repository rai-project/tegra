all: generate

fmt:
	go fmt ./...

install-deps:
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/elazarl/go-bindata-assetfs/...

glide-install:
	glide install --force

logrus-fix:
	rm -fr vendor/github.com/Sirupsen
	find vendor -type f -exec sed -i 's/Sirupsen/sirupsen/g' {} +

generate: clean generate0

generate0:
	go-bindata -nomemcopy -pkg tegra -o tegra_assets_static.go -ignore=.DS_Store  -ignore=README.md _fixtures/...

clean-models:
	rm -fr tegra_assets_static.go

clean: clean-models

travis: install-deps glide-install logrus-fix generate
	echo "building..."
	go build
