FROM busybox
LABEL MAINTAINER="jiazhu3@cisco.com"
ADD dist/ginsupercmd_linux /
CMD ["/ginsupercmd_linux"]

#docker run -it --rm -p 9180:8180 -e GIN_MODE='release'  supercmd


