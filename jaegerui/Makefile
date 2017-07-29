all: jaeger_ui_resources.go

fmt:
	go fmt ./...
	go vet ./...

jaeger_ui_resources.go: jaeger-ui-build
	mkdir -p $(dir $@)
	go-bindata -nomemcopy -o $@ -prefix ${PWD}/jaeger-ui-build/build -pkg jaegerui $</build/...

checkout-ui:
	git clone --depth 1 https://github.com/uber/jaeger-ui

jaeger-ui-build:
	cd jaeger-ui && yarn install && npm run build
	rm -rf jaeger-ui-build && mkdir jaeger-ui-build
	cp -r jaeger-ui/build $</

clean-ui:
	rm -fr jaeger-ui

clean:
	rm -r jaeger_ui_resources.go
