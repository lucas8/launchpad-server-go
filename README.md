# Launchpad Server

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
