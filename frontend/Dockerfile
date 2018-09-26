FROM node:8.11.4-alpine as builder
ARG ENV=prod
COPY package.json ./
RUN yarn install --ignore-engines && mkdir /app && cp -R ./node_modules ./app
WORKDIR /app
COPY . .
RUN source .env.${ENV} && yarn run build

FROM nginx:1.15.2-alpine
COPY nginx/default.conf /etc/nginx/conf.d/
RUN rm -rf /usr/share/nginx/html/*
COPY --from=builder /app/dist /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]