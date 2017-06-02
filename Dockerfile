FROM scratch

MAINTAINER Sebastian Mandrean <sebastian.mandrean@gmail.com>

ENV PATH /
ENV HOME /

# Add binary as its own layer
COPY drone-kubernetes-helm /

# Add supporting binaries & files
COPY ["helm", "kubectl", "/"]
COPY ca-certificates.crt /etc/ssl/certs/
WORKDIR /tmp/
WORKDIR /.kube/

# Init helm
RUN ["helm", "init", "-c"]

CMD ["drone-kubernetes-helm"]
