FROM envoyproxy/envoy:v1.31-latest
COPY ./envoy.yaml /etc/envoy/envoy.yaml
COPY ./pb/ /etc/pb/
RUN chmod go+r /etc/envoy/envoy.yaml
ENTRYPOINT [ "/usr/local/bin/envoy", "-c", "/etc/envoy/envoy.yaml" ]
EXPOSE 51051
