FROM tensorflow/tensorflow

#Begin: install dependencies
RUN apt-get update && apt-get install -y --no-install-recommends git
#End: install dependencies

#Begin: install golang
ENV GOLANG_VERSION 1.14
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 08df79b46b0adf498ea9f320a0f23d6ec59e9003660b4c9c1ce8e5e2c6f823ca
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz && \
    echo "$GOLANG_DOWNLOAD_SHA256 golang.tar.gz" | sha256sum -c - && \
    tar -C /usr/local -xzf golang.tar.gz && \
    rm golang.tar.gz 
RUN export PATH=/usr/local/go/bin:$PATH
#End: install golang

#Begin: install tensorflow
ENV TENSORFLOW_LIB_GZIP libtensorflow-cpu-linux-x86_64-1.15.0.tar.gz
ENV TARGET_DIRECTORY /usr/local
RUN  curl -fsSL "https://storage.googleapis.com/tensorflow/libtensorflow/$TENSORFLOW_LIB_GZIP" -o $TENSORFLOW_LIB_GZIP && \
     tar -C $TARGET_DIRECTORY -xzf $TENSORFLOW_LIB_GZIP && \
     rm -Rf $TENSORFLOW_LIB_GZIP && \
     ldconfig 
#End: install tensorflow

COPY . /micro-image-recognizer
RUN cd micro-image-recognizer && CGO_ENABLED=1 GOOS=linux go build -o app
CMD ["./micro-image-recognizer/app"]