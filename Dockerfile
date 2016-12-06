# Docker image for the Drone build runner
#
#     CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t mandrean/drone-kubernetes-helm .

FROM alpine:3.4

MAINTAINER Sebastian Mandrean <sebastian.mandrean@gmail.com>

# Environment variables
ARG K8S_VERSION 1.4.6
ARG HELM_VERSION 2.0.2

# Install dependencies & Clean up
RUN apk --no-cache --update --repository http://dl-3.alpinelinux.org/alpine/edge/community/ add \
	curl \
	git \
&& apk --no-cache del \
	wget \
&& rm -rf /var/cache/apk/* /tmp/*

# Install kubectl & helm
RUN curl -o /usr/local/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/v$K8S_VERSION/bin/linux/amd64/kubectl \
&& chmod +x /usr/local/bin/kubectl \
&& curl https://storage.googleapis.com/kubernetes-helm/helm-v$HELM_VERSION-linux-amd64.tar.gz | tar zxvf - \
&& mv linux-amd64/helm /usr/local/bin/helm && rm -rf linux-amd64 \
&& mkdir -p ~/.kube/credentials && mkdir -p ~/.helm

ADD ./drone-kubernetes-helm /usr/local/bin/

ENTRYPOINT ["drone-kubernetes-helm"]
