#stage 0
FROM node:16.14.2 as helper

WORKDIR /frontend/

COPY package*.json ./

RUN npm install
COPY ./ ./

RUN npm run build

#stage 1
FROM nginx:latest

COPY --from=helper /frontend/build /usr/share/nginx/html

COPY --from=helper /frontend/nginx.conf /etc/nginx/nginx.conf