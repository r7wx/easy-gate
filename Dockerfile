FROM golang:1.18 AS go-builder
WORKDIR /easy-gate
COPY . .
RUN make

FROM node:17 AS web-builder
WORKDIR /easy-gate-web
COPY ./web .
RUN yarn install && yarn build

FROM nginx:1.21.6-alpine AS easy-gate
WORKDIR /usr/bin
COPY --from=go-builder ./easy-gate/build/easy-gate .
WORKDIR /etc/easy-gate
COPY ./assets/easy-gate.json .
WORKDIR /usr/share/nginx/html
COPY --from=web-builder ./easy-gate-web/build .
WORKDIR /etc/nginx/conf.d/
COPY ./assets/nginx.conf default.conf
WORKDIR /
COPY ./docker-entrypoint.sh .
RUN chmod +x ./docker-entrypoint.sh
