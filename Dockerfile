FROM fedora:30

RUN dnf install golang -y
RUN mkdir -p $HOME/go
RUN echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc
RUN source $HOME/.bashrc
RUN go env GOPATH

RUN dnf install nodejs -y
RUN npm install -g http-server