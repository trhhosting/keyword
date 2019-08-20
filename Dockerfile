FROM quay.io/spivegin/golang_dart_protoc_dev AS build-env
WORKDIR /opt/src/src/github.com/trhhosting/keyword
ADD . /opt/src/src/github.com/trhhosting/keyword

RUN apt-get -y update && apt-get -y upgrade && \
    apt-get -y install gcc && apt-get -y autoremove && apt-get -y clean &&\
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENV deploy=2619b131d7ceeab1002fbf5e67b58ab6692a7123 \
    GO111MODULE=on
RUN git config --global url."https://${deploy}@sc.tpnfc.us".insteadOf "https://sc.tpnfc.us"

RUN cd /opt/src/src/github.com/trhhosting/keyword &&\
    go mod tidy &&\
    go get ./... &&\
	go build -o build/keyword ./keyword


FROM quay.io/spivegin/tlmbasedebian
COPY --from=build-env /opt/src/src/github.com/trhhosting/keyword/build/keyword /bin/keyword
CMD ["keyword", "up"]
