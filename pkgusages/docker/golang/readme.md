1. From Creating our image using the base image golang:alpine. This image runs the alpine Linux distribution which is small and has Golang already
   installed which is perfect for our use case.

2. Env Set necessary environment variables needed for our image

3. WORKDIR, COPY, RUN WORKDIR moves to working directory /build. Copy and download dependency using go mod RUN is used
   to build the application.

4. EXPORT, CMD Using export port 3000 from inside our container to the outside since the application will listen to this
   port to work. Define a default command to execute when we run our image which is CMD [“/dist/main”].

5. Using "docker build . -t go-dock" to build image. The image is named with "go-dock".

6. Run the "go-dock" image with "docker run -p 3000:3000 go-dock"
   The flag -p is to define the port binding. Since the app inside the container is running on port 3000 then here bind
   it to the host port, this time also 3000.

**7. Access your docker, confirm it works. After you run the go-dock docker,**

% docker run -d -p 3001:3000 go-dock /main

access 127.0.0.1:3001/ping

access 127.0.0.1:3001/animal/cat You will see "

{“message”:”pong”}" as your expect. You have a full working web server already.

Using 3001:3000, it won't occupy the local 3000 port, it occupies 3001. The port 3000 is used which belong to docker, not the local.

One of Docker’s best practice is keeping the image size small, by having only the binary file.

Get docker current run container id

% docker ps

CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES

894b99118767 go-dock   "/main /main"   22 seconds ago Up 21 seconds 0.0.0.0:3001->3000/tcp pensive_ptolemy

% docker stop container-id
