FROM alpine:edge

# Install specific version of MuPDF
RUN apk add --no-cache libc6-compat
# Try to find and install the specific version
RUN apk add --no-cache mupdf=1.24.0-r0 mupdf-dev=1.24.0-r0 || \
    (echo "Specific version not available, checking available versions:" && \
     apk search -v mupdf)

COPY main ./main
COPY uploads ./uploads
RUN chmod +x /main
EXPOSE 8000
CMD ["/main"]