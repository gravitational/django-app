FROM quay.io/gravitational/debian-grande:0.0.1

RUN apt-get update && \
    apt-get install --no-install-recommends --yes \
        gcc \
        libpq-dev \
        postgresql-client \
        gettext \
        python3 \
        python3-dev \
        python3-pip && \
    apt-get clean

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY . /usr/src/app

RUN pip3 install -r requirements.txt

EXPOSE 8000

CMD ["./entrypoint.sh"]
