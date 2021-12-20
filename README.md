# Go Web Demo

This is the demo code and also can be used as a template when creating a web application in GoLang

## Installation

Installation links used in this demo are given below: 
* [Go](https://go.dev/doc/install)
* [Docker](https://docs.docker.com/get-docker)
* [kubectl](https://kubernetes.io/docs/tasks/tools/)
* [minikube](https://minikube.sigs.k8s.io/docs/start/)
* [Couchbase Server](https://www.couchbase.com/downloads)
* [Goland IDE (optional)](https://www.jetbrains.com/go/download/)

## Creating Database
* After installing Couchbase server and starting it, you should see the dashboard via [localhost:8091](http://localhost:8091). 
* You can log in with the credentials you created while installing it.
* Add a bucket named `reviews` from top right corner on the dashboard.
* Run command below from query section on the left to be able to run n1ql queries on the `reviews` bucket
```n1ql
CREATE PRIMARY INDEX ON `default`:`reviews`
```
*** 
**Reminder**: Your database credentials should match with the credentials in `resource/application.yml`
***

## Running the Server

In the main directory of the project
```bash
#build go app
go build -t main

#then simply run it
./main
```

## Dockerize
In this section, we will dockerize our go app and push it to DockerHub in order to run our deployments and create our pods in minikube.

One should have a docker hub account to be able to push a docker image to the hub. You should sign up [here](https://hub.docker.com/) to continue from here.

Once you signed up you should log in from your local computer via:
```bash
docker login
```
Then you can continue with:
```bash
# Dockerize the application
docker build -t go-demo

# tag your docker image with your username and image version
docker tag go-demo <username>/go-demo:<version something like (1.x.x)>

# now you can push your tagged image to the dockerhub
docker push <username>/go-demo:1.0.0
```

After pushing your image, you should change the image being pulled from dockerhub in `.deploy/deployment.yml`
```yaml
containers:
   - name: go-demo
     image: <username>/go-demo:<version>
```

## Deploying Application to K8S
You should have kubectl and minikube installed on your machine to continue from this part.
```bash
# start minikube
minikube start
# you should see something like
😄  minikube v1.24.0 on Darwin 11.5.2
✨  Using the docker driver based on existing profile
👍  Starting control plane node minikube in cluster minikube
🚜  Pulling base image ...
🏃  Updating the running docker "minikube" container ...
🐳  Preparing Kubernetes v1.22.3 on Docker 20.10.8 ...
🔎  Verifying Kubernetes components...
    ▪ Using image k8s.gcr.io/metrics-server/metrics-server:v0.4.2
    ▪ Using image gcr.io/k8s-minikube/storage-provisioner:v5
    ▪ Using image kubernetesui/dashboard:v2.3.1
    ▪ Using image kubernetesui/metrics-scraper:v1.0.7
🌟  Enabled addons: metrics-server, storage-provisioner, default-storageclass, dashboard
🏄  Done! kubectl is now configured to use "minikube" cluster and "default" namespace by default

# this is optional but you can enable metrics-server add on to see pod usages and other metrics to monitor yor cluster
minikube addons enable metrics-server

# you can see the cluster dashboard via
minikube dashboard
#this will open the cluster dashboard in your default browser 

# let's create a namespace for our applications to deploy
kubectl create namespace reviews

# now you can create your deployments and services. cd to go-web-demo/.deploy folder and apply command below.

# this command will create a deployment  in reviews namespace according to settings given in yaml
kubectl apply - deployment.yml -n reviews

# then create your service and horizontal pod autoscaler
kubectl apply - service.yml -n reviews 
kubectl apply - hpa.yml -n reviews

# from now on you should have an application running in local minikube cluster. 

# You can now see the url to reach your running pod via: 
minikube service go-demo-service --url -n reviews

🏃  Starting tunnel for service go-demo-service.
|-----------|-----------------|-------------|------------------------|
| NAMESPACE |      NAME       | TARGET PORT |          URL           |
|-----------|-----------------|-------------|------------------------|
| reviews   | go-demo-service |             | http://127.0.0.1:54811 |
|-----------|-----------------|-------------|------------------------|
```


After this point, you can use the application and extend it according to your own needs.

## License
[MIT](https://choosealicense.com/licenses/mit/)
