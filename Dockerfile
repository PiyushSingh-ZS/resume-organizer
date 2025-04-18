FROM alpine:edge

# Install required dependencies including MuPDF
RUN apk add --no-cache libc6-compat mupdf mupdf-dev

COPY main ./main

COPY uploads ./uploads

RUN chmod +x /main

EXPOSE 8000

CMD ["/main"]