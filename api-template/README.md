# This is an api template. 

### Getting started
1. Copy the api and name it appropriate. Use the api-name convention to keep this repo organized.
2. Update the imports in router.go - e.g. "github.com/asciiu/appa/[api-template]/controllers". The api-template import needs to be udpated to reflect your new api-name.
3. Update the docker image tag names in the Makefile for build and push. Tag the docker image according to the naming convention above - e.g. us.gcr.io/asciiu/appa/api-name:dev. You must follow the us.gcr/asciiu/appa convention as that is the gcloud docker registry for appa's clusters.
4. Update the deployment.tmpl, deployment.yml, and service.yml files under deployments/prod and deployments/stage. Follow the same convention you used for the api-name. You will also have to update the docker image name in these files. e.g. us.gcr.io/asciiu/appa/api-name:dev. This ensures that the approprite image is used during deployment.
5. Update your local config-maps using api-name as well:
```
metadata:
  name: something-config
```
Update the config map ref in deployments/prod/deployment.tmp and deployments/stage/deployment.tmp to use the appropriate config ref like so:
```
- configMapRef:
    name: admin-config
```
6. Update the module name for the api in go.mod.
7. There is also an api-volume.yml file under deployments for both stage and prod. Make sure you update that volume metadata name. Under each deployment.tmpl and deployment.yml file you will also need to update the claim name to reference the appropriate vol - i.e. claimName: plasma-api-admin-vol. This persistent volume is used during the autocert - a.k.a. let's encrypt process - to obtain an SSL cert.

### Dependencies 
* go - v1.12

### Build command
```
$ export SSH_PRIVATE_KEY="$(cat ../id_rsa)"
$ make build
$ make push
```

### Run process
Use false argument param to delineate http for local dev. Use true
for https. Refer to mainMiddlewares.go to see how this works.
```
$ ./api-name false 
```

via source
```
$ go run main.go router.go false
```

### Deploy
1. Add your env vars to 
./deployments/stage/config-map.yaml
./deployments/prod/config-map.yaml

note: make sure you name it appropriately. Refer to "Getting Started"

2. Deploy the configs to stage and prod using kubectl.
```
$ kubectl create -f ./deployments/stage/config-map.yaml
```

3. Start/Restart the service using kubectl.
```
$ make start
```
or 
```
$ make update
```
