# Multi-stage build for Svelte frontend
FROM oven/bun:1-alpine AS builder

# Build arg for API URL
ARG PUBLIC_API_URL=http://localhost:8080/query

WORKDIR /build

# Install dependencies first for better caching
COPY package.json bun.lockb ./
RUN bun install

# Copy source (excluding generated files)
COPY . .

# Build for production (skip codegen, use pre-generated types)
ENV PUBLIC_API_URL=${PUBLIC_API_URL}
RUN bun vite build

# Runtime stage - nginx alpine for serving static files
FROM nginx:1.27-alpine

# Copy nginx config
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy built files
COPY --from=builder /build/build /usr/share/nginx/html

# Use non-root user
RUN chown -R nginx:nginx /usr/share/nginx/html && \
    chown -R nginx:nginx /var/cache/nginx && \
    chown -R nginx:nginx /var/log/nginx && \
    chown -R nginx:nginx /etc/nginx/conf.d && \
    touch /var/run/nginx.pid && \
    chown -R nginx:nginx /var/run/nginx.pid

USER nginx

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost/ || exit 1

STOPSIGNAL SIGQUIT

CMD ["nginx", "-g", "daemon off;"]
