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
