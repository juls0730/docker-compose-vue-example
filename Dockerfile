# Dockerfile
FROM oven/bun:alpine

RUN mkdir /app
WORKDIR /app

COPY . /app

RUN bun install

# Run the vite dev server with the bun runtime
CMD ["bun", "--bun", "dev", "-- --host=0.0.0.0"]
