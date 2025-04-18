FROM alpine:edge

RUN apk add --no-cache libc6-compat

RUN apk add --no-cache libffi ca-certificates tzdata

COPY main ./main

COPY uploads ./uploads

RUN chmod +x /main

EXPOSE 8000

CMD ["/main"]
