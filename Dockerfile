FROM ubuntu:20.04

COPY main ./main

RUN chmod +x /main

EXPOSE 8000

CMD ["/main"]
