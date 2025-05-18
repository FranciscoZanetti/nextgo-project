# Build stage
FROM node:20-alpine as builder
WORKDIR /app

COPY package*.json ./
RUN npm install

COPY postcss.config.js tailwind.config.js ./
COPY . .
RUN npm run build

# Production stage
FROM node:20-alpine as runner
WORKDIR /app

COPY --from=builder /app/public ./public
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/node_modules ./node_modules

EXPOSE 3000
CMD ["npm", "start"]
