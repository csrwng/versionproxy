FROM centos:centos7

COPY ./versionproxy /usr/bin/versionproxy
CMD  /usr/bin/versionproxy
