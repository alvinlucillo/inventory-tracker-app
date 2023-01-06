FROM node:14-alpine as build

WORKDIR /app

# COPY ./frontend/package.json .
COPY ./frontend/package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:stable-alpine

COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]