FROM registry.fedoraproject.org/fedora-minimal:37 as build

WORKDIR /usr/lib/golang/src/internetbox2-exporter
# Not ideal as it also copies all the git objects, but it is only the build container   
COPY . .
RUN microdnf install -y golang git && go get
# http://blog.wrouesnel.com/articles/Totally%20static%20Go%20builds/
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o exporter .

FROM scratch
LABEL maintainer="Simon Krenger <simon@krenger.ch>"
LABEL description="A Prometheus Exporter for the Swisscom Internetbox 2"

WORKDIR /
USER 1001

COPY --from=build /usr/lib/golang/src/internetbox2-exporter/exporter /exporter

EXPOSE 8088

CMD ["/exporter"]
