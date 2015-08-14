#
# API Frontend Dockerfile
#

# Pull the base image
FROM golang:1.4.2
MAINTAINER Cleiton Marques <cleitonmarx@hotmail.com>

# Set GOPATH
ENV GOPATH /go


# Make directories for GoWebApp
RUN mkdir -p /go/src/github.com/cleitonmarx/GoWebApp

# Add GoWebApp files
ADD . /go/src/github.com/cleitonmarx/GoWebApp

# Define working directory
WORKDIR /go/src/github.com/cleitonmarx/GoWebApp

# Restore Dependencies and Install Application
RUN \
	cd /go/src/github.com/cleitonmarx/GoWebApp && \
	go get github.com/tools/godep && \
	godep restore && \
	go install

# Define default command
CMD ["/go/bin/GoWebApp"]

# Expose Ports
EXPOSE 3333
