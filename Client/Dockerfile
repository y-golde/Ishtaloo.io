FROM node:latest AS build
COPY ./ /app
WORKDIR /app
RUN npm install && npm i -g rollup && npm run build

FROM nginx:latest
RUN mkdir /app
COPY --from=build /app/public /app
COPY nginx.conf /etc/nginx/nginx.conf