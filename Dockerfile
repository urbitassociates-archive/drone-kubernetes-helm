# Docker image for the Drone build runner
#
#     CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t urbit/drone-kubernetes-helm .

FROM alpine:3.8

# Environment variables
ARG K8S_VERSION=1.9.3
ARG HELM_VERSION=2.9.1
ARG KUBE_PATH="/root/.kube/"

ADD ./config/kubeconfig $KUBE_PATH/kubeconfig

ADD . /root/go/src/github.com/urbitassociates/drone-kubernetes-helm

# Install dependencies & Clean up
RUN apk --no-cache --update --repository http://dl-3.alpinelinux.org/alpine/3.8/community/ add \
	curl \
	git \
	musl-dev \
	go \
	glide \
&& apk --no-cache del \
	wget \
&& rm -rf /var/cache/apk/* /tmp/*

# Install Go dependencies
WORKDIR /root/go/src/github.com/urbitassociates/drone-kubernetes-helm
RUN glide install

# Build
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN go build

# Install kubectl & helm for the runtime
RUN curl -#SL -o /usr/local/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/v$K8S_VERSION/bin/linux/amd64/kubectl \
&& chmod +x /usr/local/bin/kubectl \
&& curl -#SL https://storage.googleapis.com/kubernetes-helm/helm-v$HELM_VERSION-linux-amd64.tar.gz | tar zxvf - \
&& mv linux-amd64/helm /usr/local/bin/helm && rm -rf linux-amd64 \
&& chmod +x /usr/local/bin/helm \
&& mkdir -p ~/.kube/credentials && helm init -c

ENTRYPOINT ["./drone-kubernetes-helm"]
