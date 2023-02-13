The flag -d means that the container is running as a daemon. In this mode, the terminal does not output any logs.

% docker compose up -d

Access http://localhost:8087/, it will show the local database adminer webpage.

Connect with docker container in terminal

docker container exec -it container-name bash

% docker container exec -it my-postgres bash

Cannot connect to docker postgres yet, but it can connect with local postgres database.
