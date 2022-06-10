FROM scratch

COPY dist/puma-exporter_linux_amd64_v1/puma-exporter /puma_exporter

EXPOSE 9882

ENTRYPOINT [ "/puma_exporter" ]
