
FROM golang:latest

# it is secretly debian so:
RUN sudo apt-get update && \
    sudo apt install openssh-server

CMD ["sleep infinity"]
