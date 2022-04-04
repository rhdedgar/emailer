# begin build container definition
FROM golang:latest as build

RUN /usr/local/go/bin/go install github.com/rhdedgar/emailer@master


# begin run container definition
FROM registry.access.redhat.com/ubi8/ubi-minimal as run

ADD scripts/ /usr/local/bin/
ADD email_templates/ /usr/local/bin/email_templates

COPY --from=build /go/bin/emailer /usr/local/bin

CMD /usr/local/bin/start.sh
