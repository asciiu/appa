build:
	dep ensure -update github.com/asciiu/appa
	docker build -t dockerowy/appa-user-service:dev -f ../user-service/Dockerfile .. 

prodo: 
	protoc --proto_path=${GOPATH}/src/github.com/asciiu/appa/user-service --micro_out=. --go_out=. proto/user/user.proto 

push:
	docker push dockerowy/appa-user-service:dev 

update:
	dep ensure -update github.com/asciiu/appa

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