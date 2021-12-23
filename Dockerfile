# FROM golang:latest
# ENV GO111MODULE=off

# RUN mkdir /app
# WORKDIR /app
# COPY . .

# RUN go get github.com/vmihailenco/tagparser
# RUN go get github.com/go-pg/pg
# RUN go get github.com/go-pg/pg/orm
# RUN go get github.com/gorilla/mux

# RUN go test -v .

# Specify the base image for the go app.
FROM golang:latest
ENV GO111MODULE=off
# Specify that we now need to execute any commands in this directory.
WORKDIR /app
# Copy everything from this project into the filesystem of the container.
COPY . .
# Obtain the package needed to run code. Alternatively use GO Modules. 
RUN go get -u github.com/lib/pq
RUN go get github.com/gorilla/mux
RUN go get github.com/azbshiri/common/db
RUN go get github.com/pquerna/ffjson/ffjson
# Compile the binary exe for our app.
# RUN go build -o main .