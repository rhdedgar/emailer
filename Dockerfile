# begin build container definition
#FROM registry.access.redhat.com/ubi8/ubi-minimal as build
FROM golang:latest

# Install clamav
#RUN microdnf install -y golang

#ENV GOBIN=/bin \
#    GOPATH=/go

# install clam-update
RUN /usr/bin/go install github.com/rhdedgar/emailer@master


# begin run container definition
FROM registry.access.redhat.com/ubi8/ubi-minimal as run

ADD scripts/ /usr/local/bin/

COPY --from=build /bin/emailer /usr/local/bin

CMD /usr/local/bin/start.sh