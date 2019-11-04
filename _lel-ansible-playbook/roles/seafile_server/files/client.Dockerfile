FROM ubuntu:bionic
ARG USER_ID

RUN echo -------------> User_id: ${USER_ID}
RUN apt update && apt install -y software-properties-common curl \
    && echo "deb http://ppa.launchpad.net/seafile/seafile-client/ubuntu bionic main" >>  /etc/apt/source.list \
    && echo "deb-src http://ppa.launchpad.net/seafile/seafile-client/ubuntu bionic main" >>  /etc/apt/source.list

RUN curl -o - "https://keyserver.ubuntu.com/pks/lookup?op=get&search=0x327CCA0BCDC4BC61" | apt-key add - \
    && apt update -y \
    && apt install -y seafile-gui seafile-cli procps curl grep \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /seafile-client

RUN useradd -u ${USER_ID} -U -d /seafile-client -s /bin/bash seafile \
    && usermod -G users seafile \
    && chown -R seafile:seafile  /seafile-client  \
    && su - seafile -c "seaf-cli init -d /seafile-client"
USER seafile
VOLUME ["/seafile-client"]
CMD ["seafile-applet"]
