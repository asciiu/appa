build:
	docker build -t dockerowy/appa-socket:dev .

push:
	docker push dockerowy/appa-socket:dev 

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