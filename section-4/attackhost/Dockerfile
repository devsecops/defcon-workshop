FROM kalilinux/kali-linux-docker
MAINTAINER Anshuman Bhartiya anshuman.bhartiya@gmail.com

RUN echo "deb http://http.kali.org/kali kali-rolling main contrib non-free" > /etc/apt/sources.list && \
echo "deb-src http://http.kali.org/kali kali-rolling main contrib non-free" >> /etc/apt/sources.list
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get -y update && apt-get -y dist-upgrade && apt-get clean
RUN apt-get install -y git python-pip

RUN mkdir /data
WORKDIR /data

RUN git clone https://github.com/joaomatosf/jexboss.git
WORKDIR /data/jexboss
RUN pip install -r requires.txt

ENTRYPOINT while :; do read; done