FROM node:14.15

ENV LANG C.UTF-8

RUN mkdir /app
WORKDIR /app

ADD . /app

WORKDIR /app/FE
RUN npm i -g yarn --force

ARG NODE_ENV
ENV NODE_ENV ${NODE_ENV:-development}

RUN yarn install
RUN yarn run build

CMD NUXT_HOST=0.0.0.0 yarn run start
