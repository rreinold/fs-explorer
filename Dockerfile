FROM alpine:latest
COPY fs-explorer /bin/fs-explorer
COPY ./foo /foo
WORKDIR /foo
CMD ["/bin/fs-explorer"]
