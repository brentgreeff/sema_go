FROM golang:1.14

ENV GO111MODULE=on
RUN go get -u github.com/beego/bee

ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/mathapp

# ARG GROUP_ID
# ARG USER_ID

# RUN groupadd --gid $GROUP_ID app && useradd -m -l --uid $USER_ID --gid $GROUP_ID $APP_USER
# RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
RUN mkdir -p $APP_HOME

# USER $APP_USER
WORKDIR $APP_HOME
EXPOSE 8010
CMD ["bee", "run"]
