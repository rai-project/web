all: swagger_ui_resources.go

fmt:
	go fmt ./...
	go vet ./...

swagger_ui_resources.go: swagger-ui
	mkdir -p $(dir $@)
	go-bindata -nomemcopy -o $@ -prefix ${PWD}/swagger-ui -pkg swaggerui $</...

swagger-ui:
	curl -sSL https://github.com/swagger-api/swagger-ui/archive/v3.0.14.zip > /tmp/swagger-ui.zip
	mkdir -p $@
	unzip /tmp/swagger-ui.zip swagger-ui-*/dist/* -d $@
	mv swagger-ui/swagger-ui-*/dist/* swagger-ui/
	rm -r swagger-ui/swagger-ui-*

clean:
	rm -r swagger_ui_resources.go
