FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
label maintainer "Raj Durvasula <raj.durvasula@gmail.com>"

ENV OPERATOR=/usr/local/bin/rd-tomcat-operator \
    USER_UID=1001 \
    USER_NAME=rd-tomcat-operator

# install operator binary
COPY build/_output/bin/rd-tomcat-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
