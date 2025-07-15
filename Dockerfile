FROM ubuntu:jammy

RUN cd /aas

CMD ["/bin/bash", "build.sh"]
