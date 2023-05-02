
FROM node:16-alpine

# Install system dependencies
RUN echo "Installing system dependencies.." && \
      apk update && \
      apk add bash gcc make musl-dev openssl upx && \
      echo "Cleanups.." && \
      rm -rf /var/cache/apk/*

WORKDIR /infra/vault

COPY . .

RUN make install

COPY run.sh /

CMD ["/run.sh"]
