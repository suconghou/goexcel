FROM scratch
LABEL maintainer="suconghou@gmail.com"
ADD goexcel /
ENTRYPOINT ["/goexcel"]
EXPOSE 6060
