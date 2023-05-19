# use golang 1.16 as base image
FROM golang:1.19.5

# set working directory
WORKDIR /PA2/microservice/auth-service

# copy go.mod and go.sum to container
COPY go.mod .
COPY go.sum .

# download and install go modules
RUN go mod download

# copy source code to container
COPY . .

# build the binary executable file
RUN go build -o student-service .

# expose port 8080 for the application
EXPOSE 8080

# command to start the application
CMD ["./student-service"]
