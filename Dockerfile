# begin build container definition
FROM registry.access.redhat.com/ubi8/ubi-minimal as build

# Install clamav
RUN microdnf install -y golang

ENV GOBIN=/bin \
    GOPATH=/go

# install clam-update
RUN /usr/bin/go install github.com/rhdedgar/email-confirm@master


# begin run container definition
FROM registry.access.redhat.com/ubi8/ubi-minimal as run

ADD scripts/ /usr/local/bin/

COPY --from=build /bin/email-confirm /usr/local/bin

CMD /usr/local/bin/start.sh