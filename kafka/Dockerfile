FROM hub.fshome.net/golang/golang:1.12.7-buster
# Change the base image for your build

MAINTAINER Genesis

ARG HTTP_PROXY
ARG HTTPS_PROXY
ARG GO_PROXY
ARG GO_MODULE
ARG SRC_STORE

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV http_proxy $HTTP_PROXY
ENV https_proxy $HTTPS_PROXY
ENV no_proxy *.fshome.net

ENV DEBIAN_FRONTEND noninteractive

# Turn on module mode, refer to 'go help modules' for more details
ENV GOPROXY ${GO_PROXY:-https://mirrors.aliyun.com/goproxy/}

ENV LCOAL_SRC_STORE ${SRC_STORE:-http://gitlab.fshome.net/go/dependencies/raw/master/}
ENV ZMQ_VERSION 4.3.2
ENV CZMQ_VERSION 4.2.0
ENV LIBRDKAFKA_VERSION 1.1.0

# Avaliable value is off, on and auto (default is off)
ENV GO111MODULE ${GO_MODULE:-off}

ENV GOLANG_VERSION 1.12.7

# Install build tools
RUN curl -Lk -o debian.list ${LCOAL_SRC_STORE}/debian-10-ustc.list && \
    mv /etc/apt/sources.list /etc/apt/sources.list.backup && \
    mv debian.list /etc/apt/sources.list && \
    apt-get update && apt-get install -y --no-install-recommends \
       # g++ \
       # gcc \
       # libc6-dev \
       # make \
       # pkg-config \   These packages are already the latest version in golang:1.12.7-buster
       apt-utils \
       upx-ucl \
       zip \
       libmagick++-dev && \
    # Install support libraries for librdkafka
    apt-get install -y --no-install-recommends \
      libpthread-stubs0-dev \
      libssl-dev \
      # libsasl2-dev \
      libzstd-dev \
      liblz4-dev \
      zlib1g-dev \
      rapidjson-dev && \
    # Install support libraries for ZeroMQ
    apt-get install -y --no-install-recommends \
      libsodium-dev \
      libsodium23 \
      libunwind-dev \
      libunwind8 && \
     # libzmq5-dev \
     # libzmq5 
     # libczmq-dev \
     # libczmq4
    apt-get clean

RUN rm -rf /var/lib/apt/lists/*

RUN curl -Lk -o /root/zeromq-${ZMQ_VERSION}.tar.gz ${LCOAL_SRC_STORE}/zeromq-${ZMQ_VERSION}.tar.gz && \
    tar -xzf /root/zeromq-${ZMQ_VERSION}.tar.gz -C /root && \
    cd /root/zeromq-${ZMQ_VERSION} && \
    ./configure && \
    make check && \
    make install && \
    rm -f /root/zeromq-${ZMQ_VERSION}.tar.gz

RUN curl -Lk -o /root/czmq-${CZMQ_VERSION}.tar.gz ${LCOAL_SRC_STORE}/czmq-${CZMQ_VERSION}.tar.gz && \
    tar -xzf /root/czmq-${CZMQ_VERSION}.tar.gz -C /root && \
    cd /root/czmq-${CZMQ_VERSION} && \
    ./configure --with-libsodium && \
    make check && \
    make install && \
    ldconfig && \
    rm -f /root/czmq-${CZMQ_VERSION}.tar.gz

# Compile librdkafka from source
RUN curl -Lk -o /root/librdkafka-${LIBRDKAFKA_VERSION}.tar.gz ${LCOAL_SRC_STORE}/librdkafka-${LIBRDKAFKA_VERSION}.tar.gz && \
    tar -xzf /root/librdkafka-${LIBRDKAFKA_VERSION}.tar.gz -C /root && \
    cd /root/librdkafka-${LIBRDKAFKA_VERSION} && \
    ./configure && \
    make && \
    make install && \
    rm -f /root/librdkafka-${LIBRDKAFKA_VERSION}.tar.gz

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

RUN mkdir -p $GOPATH/src/golang.org/x && \
    cd $GOPATH/src/golang.org/x && \
    git clone https://go.googlesource.com/sys && \
    git clone https://go.googlesource.com/crypto && \
    git clone https://go.googlesource.com/net && \
    git clone https://go.googlesource.com/text


