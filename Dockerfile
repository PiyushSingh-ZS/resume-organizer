FROM ubuntu:20.04

RUN apk add --no-cache tzdata ca-certificates

COPY main ./main

RUN chmod +x /main

EXPOSE 8000

CMD ["/main"]
