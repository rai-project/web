all: jaeger_ui_resources.go

fmt:
	go fmt ./...
	go vet ./...

jaeger_ui_resources.go: jaeger-ui-build
	mkdir -p $(dir $@)
	go generate

checkout-ui:
	git clone --depth 1 --branch feature/mlmodelscope https://github.com/rai-project/jaeger-ui

jaeger-ui-build:
	cd jaeger-ui && yarn install && yarn build
	rm -rf jaeger-ui-build && mkdir jaeger-ui-build
	cp -r jaeger-ui/packages/jaeger-ui/build jaeger-ui-build/

clean-ui:
	rm -fr jaeger-ui

clean:
	rm -r jaeger_ui_resources.go
