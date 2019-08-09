FROM docker.neg/dfis/golang:1.12.7-buster
# Change the base image for your build

MAINTAINER Genesis

ARG HTTP_PROXY
ARG HTTPS_PROXY
ARG GO_PROXY
ARG GO_MODULE
ARG RESOURCE_SVR

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV http_proxy $HTTP_PROXY
ENV https_proxy $HTTPS_PROXY
ENV no_proxy git.newegg.org,st01nbx01

ENV RESOURCE_DEPOT ${RESOURCE_SVR:-http://st01nbx01:8777}

# Turn on module mode, refer to 'go help modules' for more details
ENV GOPROXY ${GO_PROXY:-https://mirrors.aliyun.com/goproxy/}

# Avaliable value is off, on and auto (default is off)
ENV GO111MODULE ${GO_MODULE:-off}

ENV GOLANG_VERSION 1.12.7

# Install build tools
RUN curl -Lk -o debian-10-buster.list ${RESOURCE_DEPOT}/linux-repo/debian-10-buster-aliyun.list && \
    mv /etc/apt/sources.list /etc/apt/sources.list.backup && \
    mv debian-10-buster.list /etc/apt/sources.list && \
    DEBIAN_FRONTEND=noninteractive apt-get update -y -q && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y -q --force-yes --no-install-recommends apt-utils && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y -q --force-yes --no-install-recommends \
       build-essential \
       git-core libtool \
       autotools-dev \
       autoconf \
       automake \
       pkg-config \
       unzip \
       uuid-dev \
       libkrb5-dev \
       cmake \
       upx-ucl \
       zip \
       libmagick++-dev && \
    # Install support libraries for librdkafka
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
      libpthread-stubs0-dev \
      libssl-dev \
      # libsasl2-dev \
      libzstd-dev \
      liblz4-dev \
      zlib1g-dev \
      libunwind-dev \
      libunwind8 \
      rapidjson-dev && \
    # Install support libraries for ZeroMQ
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
      libsodium-dev \
      libsodium23 \
      liblzma-dev
     # libzmq5-dev \
     # libzmq5 
     # libczmq-dev \
     # libczmq4

RUN rm -rf /var/lib/apt/lists/*

ENV ZMQ_VERSION 4.3.2
RUN curl -Lk -o /root/zeromq-${ZMQ_VERSION}.tar.gz ${RESOURCE_DEPOT}/static/zeromq-${ZMQ_VERSION}.tar.gz && \
    tar -xzf /root/zeromq-${ZMQ_VERSION}.tar.gz -C /root && \
    cd /root/zeromq-${ZMQ_VERSION} && \
    ./configure --without-docs --enable-static --enable-shared && \
    LDFLAGS=-lstdc++ make && \
    make install && \
    ldconfig && \
    rm -f /root/zeromq-${ZMQ_VERSION}.tar.gz

ENV CZMQ_VERSION 4.2.0
RUN curl -Lk -o /root/czmq-${CZMQ_VERSION}.tar.gz ${RESOURCE_DEPOT}/static/czmq-${CZMQ_VERSION}.tar.gz && \
    tar -xzf /root/czmq-${CZMQ_VERSION}.tar.gz -C /root && \
    cd /root/czmq-${CZMQ_VERSION} && \
    ./configure --without-docs --enable-static --enable-shared && \
    make && \
    make install && \
    ldconfig && \
    rm -f /root/czmq-${CZMQ_VERSION}.tar.gz
    
ENV LIBRDKAFKA_VERSION 1.1.0
# Compile librdkafka from source
RUN curl -Lk -o /root/librdkafka-${LIBRDKAFKA_VERSION}.tar.gz ${RESOURCE_DEPOT}/static/librdkafka-${LIBRDKAFKA_VERSION}.tar.gz && \
    tar -xzf /root/librdkafka-${LIBRDKAFKA_VERSION}.tar.gz -C /root && \
    cd /root/librdkafka-${LIBRDKAFKA_VERSION} && \
    ./configure --enable-static && \
    make && \
    make install && \
    ldconfig && \
    rm -f /root/librdkafka-${LIBRDKAFKA_VERSION}.tar.gz


RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

RUN mkdir -p $GOPATH/src/golang.org/x && \
    cd $GOPATH/src/golang.org/x && \
    git clone https://go.googlesource.com/sys && \
    git clone https://go.googlesource.com/crypto && \
    git clone https://go.googlesource.com/net && \
    git clone https://go.googlesource.com/text



