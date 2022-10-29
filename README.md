# Billing Application

This project is a simple backend service that receives billing events from different sources. It saves the data in a MongoDB server.

## Deployment Instruction
### Spinning Up Services cluster
We will fetch the latest MongoDB, zookeeper and kafka docker image and run it in a docker. It can be done with the following commands
```shell
docker docker compose up -d
```

### Running The Application
Now we will build the application and run it in another container
```shell
docker build -f build/Dockerfile . -t billing_application:latest
docker run -d -p 8080:8080 billing_application
```