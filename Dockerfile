FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y mupdf mupdf-tools libffi8 tzdata ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY main /main
COPY uploads /uploads

RUN chmod +x /main
EXPOSE 8000

CMD ["/main"]