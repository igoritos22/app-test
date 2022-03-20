## App Test Go App

This is a simple web application in go lang.

The main propose of this app is give support to devops  labs like deploy apps in kubernetes, best pratices in contarnization, helming apps and others

### Building the app
After that you make your Dockerfile, run:

```console
docker build --tag app-test:<your_tag_version> .
```

### Pushing the docker image to Docker Registry

You can have others container registry, here we are using Docker Hub Registry. To push the docker image to Docker Hub Registry, run:

```console
docker login -u <your_dockerhub_username> && \
docker tag app-test:<your_tag_version> igoritosousa22/app-test:<tag_version> && \
docker push igoritosousa22/app-test:<tag_version>
```
### Deploy on Kubernetes

To deploy this app on Kubernetes, first create a manifest deploy passing the URL container registry on the image entry:
```yaml
#path: infra/app-test/yaml/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-app-test
  labels:
    app: go-app-test
spec:
  replicas: 1
  selector:
    matchLabels:
      app: go-app-test
  template:
    metadata:
      labels:
        app: go-app-test
    spec:
      containers:
        - name: go-app-test
          image: docker.io/igoritosousa22/app-test:1.0.0
          imagePullPolicy: Always
          ports:
            - containerPort: 8080
              protocol: TCP
```

### Create a Helm Chart for app-test

First, you must create the app-test files and directories. Then, run:
```console
helm create app-test
```

So, the directory app-test will be created with some other files, Let's see one of them:

  * charts - where the dependencies stored
  * templates - is where the templates for Deployments, Services, ConfigMap, etc are stored.
  * .helmignore - files that will bem ignore in process. Like dockerignore or gitignore file.
  * Chart.yaml - is where are defined informations about Chart
  * values.yaml - defines the values that will be used by templates for yor chart.

Let's run the chart:

```console
helm install app-test go/app-test/
```

If all works, you must see the output like these:
```console
NAME: app-test
LAST DEPLOYED: Sun Mar 20 16:17:01 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
```
### Using values.yaml

The helm charts allow you customize templates to apply your deployments. So, we will pass the values of our deployment in the values.yaml file. Firts edit the template/deployment.yaml with the placeholder values: **{{.Value}}**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Values.name}}
  labels:
    app: {{.Values.name}}
spec:
  replicas: {{.Values.deployment.replicas}}
  selector:
    matchLabels:
      app: {{.Values.name}}
  template:
    metadata:
      labels:
        app: {{.Values.name}}
    spec:
      containers:
        - name: {{.Values.name}}
          image: {{.Values.deployment.image}}:{{.Values.deployment.tag}}
          imagePullPolicy: Always
          ports:
            - containerPort: {{.Values.deployment.containerPort}}
              protocol: TCP
```

So, define the deployment values in the values.yaml:

```yaml
name: app-test
deployment:
  replicas: 2
  image: docker.io/igoritosousa22/app-test
  tag: 2.0.0
  containerPort: 8080
```

To see what will be deployed, run this command:

```console
helm template app-test go/app-test/ --debug
```


Apply the release running the command above

```console
helm install app-test go/app-test/
```

And, If all it works, you should see something like these:
```console
NAME: app-test
LAST DEPLOYED: Sun Mar 20 16:40:56 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
```

You can verify if the deployment is running on cluster:

```console
$ kubectl get pods 

NAME                        READY   STATUS    RESTARTS   AGE
app-test-6b5d567b89-gsl6j   1/1     Running   0          14m
app-test-6b5d567b89-xzhwf   1/1     Running   0          14m
```

### Upgrade Chart

So, Let's create a service to expose our deployment on templates/service.yaml:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: svc-{{.Values.name}}
spec:
  ports:
    - name: http
      port: {{.Values.service.port}}
  selector:
    app: {{.Values.name}}
```

Then, run:
```console
helm template app-test go/app-test/ --debug
```

If the output works fine and don't show any errors, upgrade your release
```console
helm upgrade app-test go/app-test/
```

Now you should see that que value of REVISION change:

```console
NAME: app-test
LAST DEPLOYED: Sun Mar 20 17:05:21 2022
NAMESPACE: default
STATUS: deployed
REVISION: 2
TEST SUITE: None
````

Let's see if the service was created in the cluster:
```console
$ kubectl get svc
NAME           TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
svc-app-test   ClusterIP   172.20.190.114   <none>        80/TCP    2m1s
```

Expose the service to test locally the app-test
```console
kubectl port-forward svc/svc-app-test 8080:8080
```

Test the app-test with a HTTP request
```console
curl -v http://localhost:8080/v2

(...)
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host localhost left intact
This apps works on  - Version 2!
```

To clean-up the release, just run:
```console
helm uninstall app-test go/app-test/
```

### References
https://medium.com/geekculture/helm-in-kubernetes-part-2-how-to-create-a-simple-helm-chart-af899fc2741d


