start-mock:
	docker-compose up -d mock_server
	docker-compose up waitfor

start-mock-verbose:
	docker-compose up mock_server

stop:
	docker-compose stop

integration: start-mock
	docker-compose up integration
