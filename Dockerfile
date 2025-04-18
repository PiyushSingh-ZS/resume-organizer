FROM alpine:edge

# Install glibc (this requires adding the glibc package)
RUN apk add --no-cache libc6-compat

COPY main ./main

RUN chmod +x /main

EXPOSE 8000

RUN ls -la

CMD ["/bin/sh", "-c", "/main & tail -f /dev/null"]