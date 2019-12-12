# Launchpad Server

[![CircleCI](https://circleci.com/gh/getlaunchpad/launchpad-server.svg?style=svg)](https://circleci.com/gh/getlaunchpad/launchpad-server)

> Server still in very early phases of development

### Helpful Commands

##### Pushing code via GitHub Packages

If you need help getting your creds in order check out [this gist](https://gist.github.com/LucasStettner/66b2108d0fd9663f2c09db5556f69d39)

```shell
# Build docker container
$ docker build -t  docker.pkg.github.com/getlaunchpad/launchpad-server/launchpad:latest .

# Publish to Github Packages
$ docker push docker.pkg.github.com/getlaunchpad/launchpad-server/launchpad:latest
```

##### Deploying kubernetes locally

```shell
# Started k8 process
$ minikube start

# Apply k8 config
$ kubectl apply -f k8s-deployment.yml

# Open node port
$ minikube service launchpad-service
```
