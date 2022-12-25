1. From Creating our image using the base image golang:alpine. This image runs the alpine Linux distribution which is
   small in size and has Golang already installed which is perfect for our use case.

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
% docker run -d -p 3000:3000 go-dock /main access 127.0.0.1:3000/ping access 127.0.0.1:3000/animal/cat You will see "
{“message”:”pong”}" as your expect. You have a full working web server already.

One of Docker’s best practice is keeping the image size small, by having only the binary file.
