
# Dockerfile for Express.js (Development)

FROM node:18-slim

WORKDIR /usr/src/app/

COPY package.json package-lock.json /usr/src/app/

RUN npm install

COPY . .

EXPOSE 3000

CMD ["sh", "-c", "npm start --host=0.0.0.0 --port=3000"]
