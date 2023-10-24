FROM golang:1.21.1

RUN mkdir /go-sekeleton
WORKDIR /go-skeleton

COPY go.mod go.sum ./
RUN go mod download

# Copy every thing other than go.mod & go.sum
COPY . ./

#RUN rm -Rf ./.env

# Build client & run
RUN go build -o gs-c ./client
RUN chmod +x ./gs-c

# Build server & run
RUN go build -o gs-gd ./go-dummy
RUN chmod +x ./gs-gd

# Make the script executable
RUN chmod +x ./run.sh

CMD ["./run.sh"]