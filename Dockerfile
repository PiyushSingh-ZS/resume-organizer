FROM alpine:edge

COPY main ./main

RUN chmod +x /main

EXPOSE 8000

RUN ls -la

CMD ["/main"]
