## Syborg: Discord Bot

Syborg is a pluggable discord bot with a Web UI for configuration and support for persistenting state with a built-in database.

# Usage

### Configuration
Use the example below as a template to create ./config.json
```{
	"bot_token": "Bot <TOKEN_GOES_HERE>",
	"owner_id": "BotMcGee#1337",
	"use_sharding": false,
	"shard_id": 0,
	"shard_count": 1
}
```

### Local Docker
1. Pull container image: `docker pull tzakrajs/syborg`
1. Run container image with config attached: `docker run tzakrajs/syborg -v ./config.json:/config.json -p 8080:80`

### K8s
1. Create configmap for the syborg config file `kubectl create configmap syborg-config --from-file=config.json`
1. Apply manifest to deploy syborg deployment `kubectl apply -f ./k8s-deployment.yaml`