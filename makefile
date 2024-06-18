.PHONY: build run stop clear

build:
	docker build -t myapp .

run:
	docker run -it -p 8080:8080 --name myapp-container myapp

stop:
	docker stop myapp-container 
	docker rm myapp-container 

clear: stop
	docker rmi myapp 
