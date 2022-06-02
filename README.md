# Syborg: Discord Bot

Syborg is a discord bot based on GoMusicBot, containerized, k8s-ready and lightweight

## Usage

### Configuration
Use the example below as a template to create ./config.json
```
{
	"bot_token": "Bot <TOKEN_GOES_HERE>",
	"owner_id": "YourName#1337",
	"use_sharding": false,
	"shard_id": 0,
	"shard_count": 1
}
```

### Local Docker
1. Pull container image: `docker pull Cloud-Fortress/syborg`
1. Run container image with config attached: `docker run Cloud-Fortress/syborg -v ./config.json:/config.json -p 8080:80`

### K8s
1. Create configmap for the syborg config file `kubectl create configmap syborg-config --from-file=config.json`
1. Apply manifest to deploy syborg deployment `kubectl apply -f ./k8s-deployment.yaml`
