
# Dockerfile for Express.js (Production)

FROM node:18-slim as builder

WORKDIR /usr/src/app/

COPY package.json package-lock.json /usr/src/app/
RUN npm install --production

COPY . .

# Final production image
FROM node:18-slim

WORKDIR /usr/src/app/

COPY --from=builder /usr/src/app/ /usr/src/app/

EXPOSE 3000

CMD ["sh", "-c", "npm start --host=0.0.0.0 --port=3000"]
