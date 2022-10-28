# Billing Application

This project is a simple backend service that receives billing events from different sources. It saves the data in a MongoDB server.

## Deployment Instruction
### Spinning Up MongoDB cluster
We will fetch the latest MongoDB docker image and run it in a docker. It can be done with the following commands
```shell
docker pull mongo
docker run -d -p 27017:27017 mongo
```

### Running The Application
Now we will build the application and run it in another container
```shell
docker build -f build/Dockerfile . -t billing_application:latest
docker run -d -p 8080:8080 billing_application
```