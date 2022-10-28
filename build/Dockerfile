FROM golang:1.19.2-alpine AS build
LABEL maintainer="anshumusaddi@gmail.com"


ENV SRC_DIR=/go/src/github.com/anshumusaddi/billing_application
ENV SVC=billing_application
RUN apk update && apk add --no-cache git openssh-client && rm -rf /var/cache/apk/*

# To import private go repo
ENV GOPRIVATE github.com/anshumusaddi/*
RUN git config --global url."git@github.com:anshumusaddi".insteadOf "https://github.com/anshumusaddi"
RUN mkdir ~/.ssh
RUN ssh-keyscan github.com >> ~/.ssh/known_hosts


WORKDIR ${SRC_DIR}
COPY .. ${SRC_DIR}

RUN echo $SRC_DIR
RUN echo $SVC
RUN GIT_VERSION=$(git --no-pager log -1 --pretty=format:%H)
RUN GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
RUN --mount=type=ssh,id=github GO111MODULE=on CGO_ENABLED=0 go build -o ${SVC} -ldflags "-X main.GitBranch=${GIT_BRANCH} -X main.GitVersion=${GIT_VERSION}" ./cmd

FROM golang:1.19.2-alpine
LABEL maintainer="anshumusaddi@gmail.com"

ENV HOME=/app
ENV SRC_DIR=/go/src/github.com/anshumusaddi/billing_application
RUN mkdir -p $HOME
WORKDIR $HOME

ENV SVC=billing_application

COPY --from=build ${SRC_DIR}/${SVC} ./
COPY --from=build ${SRC_DIR}/config ./config
COPY --from=build ${SRC_DIR}/build/entrypoint.sh ./entrypoint.sh

EXPOSE 8080
CMD ["sh", "entrypoint.sh" ]

