# EFISHERY Auth App

Efishery's authentication application.

- [EFISHERY Auth App](#efishery-auth-app)
  - [Usage](#usage)
  
## Usage

Install using Docker, run this command from the application root directory (restapi-service/) to build the Docker image :
```
 $ docker build -t restapi-service -f Dockerfile .
```
This will create the docker image that will be used to run the container. 

To run the container run this command :
```
 $ docker-compose -f deployments/docker-compose.yml up -d
```
This will run the container in detached mode, to run the container in non-detached mode, simply remove the `-d` flags from the command.

Run this command if you want to get into the container shell :
```
 $ docker run -it restapi-service sh
```

Run this command to stop the detached container :
```
 $ docker-compose -f deployments/docker-compose.yml down
```
