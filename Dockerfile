FROM golang:alpine as base

WORKDIR /app

RUN apk update && apk add bash inotify-tools && apk add git

ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

COPY ./go.mod /app/ 

FROM base as dev

RUN git clone https://github.com/go-delve/delve.git && \ 
    cd delve && \
    go install github.com/go-delve/delve/cmd/dlv

COPY ./ /app

RUN go mod tidy 

RUN go build -o /server -gcflags -N -gcflags -l

ENTRYPOINT sh startScript.sh

FROM base as prod

COPY . /app/

RUN go mod tidy
RUN go build -o /server

CMD ["/server"]

FROM base as test  
CMD ["go", "test", "-v", "./tests/..."]