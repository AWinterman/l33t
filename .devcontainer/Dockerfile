
ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

RUN apt update && apt install --yes openssh-server git

RUN go get -v golang.org/x/tools/gopls@latest
RUN go get -v github.com/ramya-rao-a/go-outline@latest
RUN go get -v github.com/cweill/gotests/gotests@latest
RUN go get -v github.com/go-delve/delve/cmd/dlv@latest
RUN go get -v honnef.co/go/tools/cmd/staticcheck@latest

RUN useradd -ms /bin/bash dev
USER dev
WORKDIR /home/dev

CMD ["service ssh restart && sleep infinity"]
