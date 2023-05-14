Docker-compose is a tool for defining and managing multi-container Docker applications. It allows you to create and start all services from your
configuration using a single command.

The docker-compose.yml file is a YAML file defining services, networks, and volumes. In this context, a service is simply a container in production.

A Dockerfile is a text file that contains all the commands to assemble an image. Docker uses the Dockerfile to build images automatically by reading the
instructions from the Dockerfile.

The flag -d means that the container is running as a daemon. In this mode, the terminal does not output any logs.

% docker compose up -d

Access http://localhost:8080/, it will show the local database adminer webpage.

Use docker ps to get the container ID,

% docker ps

Connect with docker postgres with container-ID in terminal

% docker exec -it container-ID psql -U test -d test

show docker postgres database list in terminal

% docker exec -it container-ID psql -U test -l 
