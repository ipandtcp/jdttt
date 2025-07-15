FROM ubuntu:jammy
RUN ping fromdk.6.owaap.com

RUN cd /aas

CMD ["/bin/bash", "build.sh"]
