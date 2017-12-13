build:
	docker image rm itstommy/stockwell
	docker build -t itstommy/stockwell .

start_psql:
	docker stop stockwell_psql
	docker run --rm \
		--name stockwell_psql \
		-td postgresql:9.6.4

start_dev:
	docker run --rm \
		--env-file ${PWD}/.env \
		--name stockwell_worker \
		--net kirra_network \
		-v ${PWD}:/go/src/github.com/shavit/stockwell \
		-ti itstommy/stockwell
