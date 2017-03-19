# Docker version Proxy
  Proxy to return semver-parseable version of Docker

## Use with 'oc cluster up'

1. Start a container with the proxy:

```
docker run --privileged --net=host -v /var/run/docker.sock:/var/run/docker.sock -d cewong/versionproxy versionproxy 127.0.0.1:2375
```

2. Set your DOCKER_HOST:

```
export DOCKER_HOST=tcp://127.0.0.1:2375
```

3. Start your cluster:

```
oc cluster up -e DOCKER_HOST=tcp://127.0.0.1:2375
```

