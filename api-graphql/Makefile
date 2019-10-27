build:
	docker build  -t dockerowy/appa-apiql:dev .

push:
	docker push dockerowy/appa-apiql:dev 

start-minikube:
	kubectl create -f ./deployments/minikube/config-map.yaml
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yaml
	kubectl create -f ./deployments/deployment.yaml
	kubectl create -f ./deployments/service.yaml

stop-minikube:
	kubectl delete -f ./deployments/deployment.yaml
	kubectl delete -f ./deployments/service.yaml
	kubectl delete -f ./deployments/minikube/config-map.yaml

restart-minikube: 
	kubectl delete -f ./deployments/minikube/config-map.yaml
	kubectl delete -f ./deployments/deployment.yaml
	kubectl delete -f ./deployments/service.yaml
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yaml
	kubectl create -f ./deployments/minikube/config-map.yaml
	kubectl create -f ./deployments/deployment.yaml
	kubectl create -f ./deployments/service.yaml