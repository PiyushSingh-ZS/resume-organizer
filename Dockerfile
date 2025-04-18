FROM debian:bookworm-slim

# Install required packages including libmupdf
RUN apt-get update && apt-get install -y \
    libmupdf-dev \
    libmujs2 \
    libharfbuzz-dev \
    libfreetype6-dev \
    libjpeg-dev \
    libopenjp2-7-dev \
    libx11-dev \
    ca-certificates \
    tzdata \
    && rm -rf /var/lib/apt/lists/*

# Copy your Go binary and assets
COPY main /main
COPY uploads /uploads

RUN chmod +x /main
EXPOSE 8000

CMD ["/main"]