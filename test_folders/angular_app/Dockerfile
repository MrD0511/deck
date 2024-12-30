
# Dockerfile for Angular (Production)

FROM node:18-slim as builder

WORKDIR /usr/src/app/

COPY package.json package-lock.json /usr/src/app/

# Install the Angular CLI globally and production dependencies
RUN npm install -g @angular/cli && npm install 

COPY . .
RUN npm run build --prod

# Final production image
FROM nginx:alpine

WORKDIR /usr/share/nginx/html

COPY --from=builder /usr/src/app/dist/angular_app/browser ./

EXPOSE 4200

CMD ["nginx", "-g", "daemon off;"]
