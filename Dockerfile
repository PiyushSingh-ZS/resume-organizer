FROM alpine:edge

RUN apk add --no-cache libc6-compat

COPY main ./main

COPY uploads ./uploads

RUN chmod +x /main

EXPOSE 8000

CMD ["/main"]
