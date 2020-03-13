# Install Protainer
### Preparation

- Get the Swarm node ID of this (manager) node and store it in an environment variable
```
export NODE_ID=$(docker info -f '{{.Swarm.NodeID}}')
```

- Create a tag in this node, so that Portainer is always deployed to the same node and uses the existing volume
```
docker node update --label-add portainer.portainer-data=true $NODE_ID
```

### Create the Docker Compose file

- create it manually, for example, using nano
```
nano portainer.yml
```

- And copy the contents inside
```
version: '3.2'

services:
  agent:
    image: portainer/agent
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - /var/lib/docker/volumes:/var/lib/docker/volumes
    networks:
      - agent_network
    deploy:
      mode: global
      placement:
        constraints: [node.platform.os == linux]

  portainer:
    image: portainer/portainer
    command: -H tcp://tasks.agent:9001 --tlsskipverify
    ports:
      - "9000:9000"
     # - "8000:8000"
    volumes:
      - portainer_data:/data
    networks:
      - agent_network
    deploy:
      mode: replicated
      replicas: 1
      placement:
        constraints: [node.role == manager]

networks:
  agent_network:
    driver: overlay
    attachable: true

volumes:
  portainer_data:
```

### Deploy it
- Deploy the stack with
```
docker stack deploy -c portainer.yml portainer
```

- Check if the stack was deployed with
```
docker stack ps portainer
```

It will output something like:
```
ID             NAME                       IMAGE                        NODE              DESIRED STATE   CURRENT STATE          ERROR   PORT
xvyasdfh56hg   portainer_agent.b282rzs5   portainer/agent:latest       dog.example.com   Running         Running 1 minute ago
j3ahasdfe0mr   portainer_portainer.1      portainer/portainer:latest   cat.example.com   Running         Running 1 minute ago
```

- You can check the Portainer logs with
```
docker service logs portainer_portainer
```

- You can update the Portainer with
```
docker service update portainer_portainer --force
```
