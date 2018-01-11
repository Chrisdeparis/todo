help:
	@echo ""
	@echo "usage: make COMMAND"
	@echo ""
	@echo "Commands:"
	@echo ""
	@echo "  gencert              Generate SSL certificates"
	@echo "  dockerBuild          Build TodoList with Docker"
	@echo ""

gencert:
	@mkdir -p etc/ssl
	@openssl genrsa -out $(shell pwd)/etc/ssl/server.key 2048
	@openssl req \
		-new \
		-x509 \
		-sha256 \
		-days 365 \
		-subj "/C=FR/ST=IDF/L=Paris/O=IPSSI/CN=localhost" \
		-key $(shell pwd)/etc/ssl/server.key \
		-out $(shell pwd)/etc/ssl/server.crt

dockerBuild:
	@make gencert
	@$(shell docker rmi -f nanoninja/todo 2> /dev/null)
	@docker build -t nanoninja/todo .

dockerRun:
	@$(shell docker rm -f nanoninja/todo 2> /dev/null)
	@docker run --name todolist -d -p 3000:3000 nanoninja/todo
