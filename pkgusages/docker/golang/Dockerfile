FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build, the working directory name you can define it by yourself,
# you also can change to /go/src/app/
WORKDIR /build

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Build the current project to a binary file called main
# The location of the binary file is /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

############################
# STEP 2 build a small image
############################
# Build a small image
FROM scratch

COPY --from=builder /dist/main /

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .
#COPY /db/data.json /db/data.json

# Command to run the binary once the container starts
ENTRYPOINT ["/main"]
