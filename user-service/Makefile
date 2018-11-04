build:
	dep ensure -update github.com/asciiu/oldiez
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/asciiu/oldiez/user-service proto/user/user.proto
	docker build -t us.gcr.io/fomo-204904/user-service:dev .

update:
	dep ensure -update github.com/asciiu/oldiez

stage:
	docker push us.gcr.io/fomo-204904/user-service:dev

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