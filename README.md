## Syborg: Discord Bot

Syborg is a pluggable discord bot with a Web UI for configuration and support for persistenting state with a built-in database.

# Usage
### Local Docker
1. Pull container image: `docker pull tzakrajs/syborg`
1. Run container image with config attached: `docker run tzakrajs/syborg -v ./config:/config -v ./data:/data -p 8080:80`

### K8s
1. Create a PVC for the database to persist through restarts call it `syborg-data`
1. Apply manifest to deploy syborg statefulset `kubectl apply ./k8s-statefulset.yaml`