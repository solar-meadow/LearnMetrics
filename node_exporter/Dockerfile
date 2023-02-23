# node-exporter

FROM ubuntu:latest
RUN apt-get update && apt-get upgrade -y && apt-get install git -y
RUN apt-get install golang -y
RUN git clone https://github.com/prometheus/node_exporter.git
RUN cd node_exporter && go build . 
EXPOSE 9100
CMD [ "./node_exporter/node_exporter" ]