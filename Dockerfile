
MAINTAINER Bitnami <containers@bitnami.com>

COPY app-code/sample-go-demo-1 /app/http-sample

USER bitnami

WORKDIR /app

EXPOSE 3000

ENTRYPOINT ["/app/http-sample"]

