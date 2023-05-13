Docker-compose is a tool for defining and managing multi-container Docker applications. It allows you to create and start all services from your
configuration using a single command.

The docker-compose.yml file is a YAML file defining services, networks, and volumes. In this context, a service is simply a container in production.

A Dockerfile is a text file that contains all the commands to assemble an image. Docker uses the Dockerfile to build images automatically by reading the
instructions from the Dockerfile.

The flag -d means that the container is running as a daemon. In this mode, the terminal does not output any logs.

% docker compose up -d

Access http://localhost:8080/, it will show the local database adminer webpage.

Connect with docker container in terminal

docker container exec -it container-name bash

% docker container exec -it my-postgres bash

Cannot connect to docker postgres yet, but it can connect with local postgres database.
