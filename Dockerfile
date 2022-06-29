
FROM golang:latest

RUN apt update && apt install --yes openssh-server git
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v github.com/ramya-rao-a/go-outline@latest
RUN go install -v github.com/cweill/gotests/gotests@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
RUN go install -v honnef.co/go/tools/cmd/staticcheck@latest

RUN useradd -ms /bin/bash dev
USER dev
WORKDIR /home/dev

CMD ["service ssh restart && sleep infinity"]
