FROM scratch

COPY bin/puma_exporter /puma_exporter

EXPOSE 9882

ENTRYPOINT [ "/puma_exporter" ]
