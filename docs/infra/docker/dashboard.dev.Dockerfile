FROM ruby:3.4-alpine

RUN apk add --no-cache \
    build-base \
    git \
    nodejs \
    npm \
    postgresql-dev \
    tzdata \
    yaml-dev

WORKDIR /app

COPY Gemfile Gemfile.lock ./

RUN bundle install
