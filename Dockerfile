FROM golang:1.11

ARG APPLICATION_DIR="/app/joyci-grpc"
ARG APPLICATION_VERSION="v0.0.1"

LABEL version=${APPLICATION_DIR}
LABEL description="This is JoyCI GRPC's Dockerfile server"

ADD . ${APPLICATION_DIR} 
WORKDIR ${APPLICATION_DIR}

RUN go build
RUN go install

EXPOSE 50051

CMD ["joyci-grpc"]