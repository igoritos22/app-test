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
docker tag app-test:<your_tag_version> igoritosousa22/app-test:<your_tag_version> && \
docker push igoritosousa22/app-test:<your_tag_version>
```
