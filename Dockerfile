FROM alpine:latest
COPY fs-explorer /bin/fs-explorer
COPY ./public /public 
WORKDIR /public
CMD ["/bin/fs-explorer"]
