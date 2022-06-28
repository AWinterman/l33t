
FROM golang:latest

# it is secretly debian so:
RUN apt-get update && \
    apt install openssh-server
    

CMD ["sleep infinity"]
