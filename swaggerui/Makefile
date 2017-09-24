all: swagger_ui_resources.go

fmt:
	go fmt ./...
	go vet ./...

swagger_ui_resources.go: swagger-ui
	mkdir -p $(dir $@)
	go-bindata -nomemcopy -o $@ -prefix ${PWD}/swagger-ui -pkg swaggerui $</...

swagger-ui:
	curl -sSL https://github.com/swagger-api/swagger-ui/archive/v3.2.2.zip > /tmp/swagger-ui.zip
	mkdir -p $@
	unzip /tmp/swagger-ui.zip swagger-ui-*/dist/* -d /tmp/swagger-ui
	mv /tmp/swagger-ui/swagger-ui-*/dist/* swagger-ui/
	rm -fr /tmp/swagger-ui
	sed -i'' -e 's/http:\/\/petstore.swagger.io\/v2\/swagger.json/\/api\/v1\/swagger.json/g' swagger-ui/swagger-ui-standalone-preset.js
	sed -i'' -e 's/http:\/\/petstore.swagger.io\/v2\/swagger.json/\/api\/v1\/swagger.json/g' swagger-ui/index.html
	rm -fr swagger-ui/*-e

clean:
	rm -r swagger_ui_resources.go
