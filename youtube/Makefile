build:
	docker build --no-cache -t dockerowy/youtube:dev --build-arg SSH_PRIVATE_KEY -f ../youtube/Dockerfile .. 

push:
	docker push dockerowy/youtube:dev 

startp:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/prod/deployment.tmpl > ./deployments/prod/deployment.yml
	kubectl create -f ./deployments/prod/deployment.yml
	kubectl create -f ./deployments/prod/service.yml

stopp:
	kubectl delete -f ./deployments/prod/deployment.yml
	kubectl delete -f ./deployments/prod/service.yml

stage: 
	kubectl delete -f ./deployments/stage/deployment.yml
	kubectl delete -f ./deployments/stage/service.yml
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/stage/deployment.tmpl > ./deployments/stage/deployment.yml
	kubectl create -f ./deployments/stage/deployment.yml
	kubectl create -f ./deployments/stage/service.yml

prod: 
	kubectl delete -f ./deployments/prod/deployment.yml
	kubectl delete -f ./deployments/prod/service.yml
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/prod/deployment.tmpl > ./deployments/prod/deployment.yml
	kubectl create -f ./deployments/prod/deployment.yml
	kubectl create -f ./deployments/prod/service.yml

updates:
	kubectl delete -f ./deployments/stage/deployment.yml
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/stage/deployment.tmpl > ./deployments/stage/deployment.yml
	kubectl create -f ./deployments/stage/deployment.yml

updatep:
	kubectl delete -f ./deployments/prod/deployment.yml
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/prod/deployment.tmpl > ./deployments/prod/deployment.yml
	kubectl create -f ./deployments/prod/deployment.yml