FROM golang AS builder

RUN apt-get update -y
RUN apt-get upgrade -y

RUN apt-get install -y locales

RUN echo "LC_ALL=en_US.UTF-8" >> /etc/environment && \
    echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen && \
    echo "LANG=en_US.UTF-8" > /etc/locale.conf && \
    locale-gen en_US.UTF-8

#RUN sudo -S DEBIAN_FRONTEND="noninteractive" apt install -y golang
RUN apt-get install -y ca-certificates && update-ca-certificates
RUN apt-get install -y make protobuf-compiler

COPY . /root/ocp-course-api

WORKDIR /root/ocp-course-api

RUN make prepare && make

FROM ubuntu
ARG service
ENV service=${service}

WORKDIR /root/
COPY --from=builder /root/ocp-course-api/${service} .
COPY --from=builder /root/ocp-course-api/swagger/${service}.swagger.json swagger/
EXPOSE 7000 7002
CMD ./${service}
