
# Dockerfile for React (Production)

FROM node:18-slim as builder

WORKDIR /usr/src/app/

COPY package.json package-lock.json /usr/src/app/
RUN npm install --production

COPY . .
RUN npm run build

# Final production image
FROM nginx:alpine

WORKDIR /usr/share/nginx/html

COPY --from=builder /usr/src/app/build .

EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]
