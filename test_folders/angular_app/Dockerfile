
# Dockerfile for Angular (Development)

FROM node:18-slim

WORKDIR /usr/src/app/

COPY package.json package-lock.json /usr/src/app/

RUN npm install

COPY . .

EXPOSE 4200

CMD ["sh", "-c", "npx ng serve --host 0.0.0.0 --port 4200"]
