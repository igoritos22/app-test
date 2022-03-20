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