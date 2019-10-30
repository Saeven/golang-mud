FROM alpine:latest

# docker build expects binaries to have already been built
ENV DIST_NAME golang-mud
ENV APP_HOME /opt/mud
WORKDIR $APP_HOME

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY ./dist $APP_HOME/
COPY ./docker/* $APP_HOME/
COPY ./resources $APP_HOME/resources
RUN chmod +x $APP_HOME/$DIST_NAME && chmod +x $APP_HOME/*.sh

ENTRYPOINT ["/opt/mud/entrypoint.sh"]
