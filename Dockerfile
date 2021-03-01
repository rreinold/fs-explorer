FROM alpine:latest
COPY fs-explorer /bin/fs-explorer
COPY ./public /public 
WORKDIR /public
CMD ["/bin/fs-explorer"]
HEALTHCHECK CMD wget -q http://0.0.0.0:3000/ -O /dev/null || exit 1
