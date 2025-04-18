FROM debian:bullseye-slim

RUN apt-get update && \
    apt-get install -y libffi8 ca-certificates tzdata && \
    rm -rf /var/lib/apt/lists/*

COPY main /main
COPY uploads /uploads
RUN chmod +x /main

EXPOSE 8000
CMD ["/main"]