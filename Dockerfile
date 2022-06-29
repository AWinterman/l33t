
FROM golang:latest

# it is secretly debian so:
RUN apt update && apt install --yes openssh-server git
    

CMD ["sleep infinity"]
