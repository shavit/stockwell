build:
	docker build -t itstommy/stockwell .

start_dev:
	docker run --rm \
		--env-file ${PWD}/.env \
		--name stockwell_1 \
		--net kirra_network \
		-v ${PWD}:/go/src/github.com/shavit/stockwell \
		-ti itstommy/stockwell
