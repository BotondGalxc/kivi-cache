FROM gcr.io/distroless/static-debian12

# Set destination for COPY
WORKDIR /app

COPY bin/kivi-server ./kivi-server

EXPOSE 5001

# Run
CMD ["./kivi-server"]