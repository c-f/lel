# --[Dockerfile for building LEL by your own]--

FROM node:12.10.0-stretch-slim AS build-frontend
COPY ./static/*js /app/
COPY ./static/.babelrc /app/
COPY ./static/src /app/src/
COPY ./static/package.json /app/
COPY ./static/package-lock.json /app/
WORKDIR /app
RUN mkdir dist \
    && npm install \
    && npm install react@^16.9.0 \
    && npm install react-dom@^16.9.0  \
    && npm run build 

# --[Builder]--
FROM golang:1.13-alpine as build-backend
WORKDIR /go/src/github.com/c-f/lel
COPY ./ .
RUN cd cmd/lel && go build -o /tmp/lel

# --[GOXC]--
# used as a template
FROM golang:1.13-alpine as goxc
WORKDIR /go/src/github.com/c-f/lel
COPY ./ . 
COPY --from=build-frontend /app/dist/app-prod.js static/dist/
RUN go get && go get github.com/laher/goxc

# --[Final LEL server]--
FROM alpine:3.10
ARG USER_ID=1000

WORKDIR /app

RUN adduser -D -u "${USER_ID}" lel \
    && mkdir -p /app/static/dist \
    && mkdir -p /project

COPY --from=build-frontend /app/dist/app-prod.js static/dist/
COPY --from=build-backend /tmp/lel /app/lel
COPY static/dist/icons /app/static/dist/icons
COPY static/dist/vendor /app/static/dist/vendor

RUN chown -R lel:lel /app && chown -R lel:lel /project
VOLUME /project
USER lel 
