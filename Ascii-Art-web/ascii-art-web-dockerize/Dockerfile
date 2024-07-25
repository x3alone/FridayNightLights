# import golang image
FROM golang:1.22.3-alpine

# MetaData
LABEL Authors="ochouari ifoukahi yhoussai"
LABEL description="This docker container will host a webserver that will print the user input in ascii-art"
LABEL source="https://learn.zone01oujda.ma/git/yhoussai/ascii-art-web-dockerize"

# working directory
WORKDIR /ascii-web

#Copy source code to working dir
COPY . .

# building the binary
RUN go build -o ascii main.go

# Expose the port
EXPOSE 8080/tcp

# command to run at container start
CMD ["./ascii"]