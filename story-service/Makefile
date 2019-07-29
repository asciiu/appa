build:
	#dep ensure -update github.com/asciiu/appa
	docker build -t dockerowy/appa-story-service:dev .

prodo: 
	protoc --proto_path=./ --micro_out=. --go_out=. proto/story/story.proto 

update: 
	dep ensure -update github.com/asciiu/appa

push:
	docker push dockerowy/appa-api:dev 

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