FROM golang:1.18.4-alpine3.16 AS builder

RUN apk --update add git make openssh

RUN addgroup -g 1001 -S builder && \
    adduser -u 1001 -S builder -G builder
USER builder:builder

ENV APP_HOME /home/builder
WORKDIR $APP_HOME

ARG SSH_PRIVATE_KEY
RUN mkdir -p ~/.ssh && echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa && chmod 0600 ~/.ssh/id_rsa \
    && git helper --global url."git@github.com:".insteadOf https://github.com/ \
    && ssh-keyscan github.com >> ~/.ssh/known_hosts

COPY --chown=builder:builder ../../.. $APP_HOME/
# build executable to ./bin/server
RUN make build-native

FROM alpine:3.15

RUN addgroup -g 1001 -S runner && \
    adduser -u 1001 -S runner -G runner

ENV BUILD_DIR /home/builder
ENV APP_HOME /home/runner/bin

RUN mkdir -p $APP_HOME

USER runner:runner
WORKDIR $APP_HOME

# copy sample .env file
COPY --from=builder $BUILD_DIR/.env.sample .env
COPY --chown=runner:runner --from=builder $BUILD_DIR/bin .

ENTRYPOINT ["$APP_HOME/{{.ProjectName}}", "http"]
CMD ["--help"]