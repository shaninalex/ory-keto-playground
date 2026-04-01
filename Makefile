migrate:
	keto migrate up --config=./config/keto.config.yml

serve:
	keto serve --config=./config/keto.config.yml

tuples:
	keto relation-tuple \
		create ./config/relation-tuples/tuples.json \
		--config=./config/keto.config.yml \
		--insecure-disable-transport-security

start:
	go run app/main.go serve --config=./config/app.config.yml