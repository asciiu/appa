build:
	docker build -t dockerowy/appa-story-service:dev .

prodo: 
	protoc --proto_path=./ --micro_out=. --go_out=. proto/story/story.proto 

push:
	docker push dockerowy/appa-story-service:dev 

start-minikube:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl create -f ./deployments/deployment.yml
	kubectl create -f ./deployments/service.yml

stop-minikube:
	kubectl delete -f ./deployments/deployment.yml
	kubectl delete -f ./deployments/service.yml

restart-minikube: 
	kubectl delete -f ./deployments/deployment.yml
	kubectl delete -f ./deployments/service.yml
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl create -f ./deployments/deployment.yml
	kubectl create -f ./deployments/service.yml