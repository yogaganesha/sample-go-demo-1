FROM alpine:latest

LABEL maintainer=Calculi \
    email=engineering@guide-rails.io

RUN echo 'OS version info...' && cat /etc/os-release \
    && apk -f -q update \
    && apk -f -q --no-cache add ca-certificates tzdata
COPY "creds/CalculiCA-cert.pem" "/usr/local/share/ca-certificates/CalculiCA-cert.crt"
RUN update-ca-certificates

RUN adduser -G root -S guiderails -u 100
USER 100
ENV PATH="${PATH}:/usr/local/bin"