build:
	dep ensure -update github.com/asciiu/oldiez
	docker build -t dockerowy/oldiez-order-service:dev -f ../order-service/Dockerfile ..

prodo: 
	protoc --proto_path=${GOPATH}/src/github.com/asciiu/oldiez/order-service --micro_out=. --go_out=. proto/order/order.proto 


push:
	docker push dockerowy/oldiez-order-service:dev 

update:
	dep ensure -update github.com/asciiu/oldiez

run:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl create -f ./deployments/deployment.yml
	kubectl create -f ./deployments/service.yml

stop:
	kubectl delete -f ./deployments/deployment.yml
	kubectl delete -f ./deployments/service.yml

restart: 
	kubectl delete -f ./deployments/deployment.yml
	kubectl delete -f ./deployments/service.yml
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl create -f ./deployments/deployment.yml
	kubectl create -f ./deployments/service.yml