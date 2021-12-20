FROM python:3.9-alpine

ENV NODE_PATH /usr/lib/node_modules/
# install nodejs
RUN apk update \
    && apk add --no-cache \
    nodejs \
    npm \
    gcc \
    libc-dev \
    git

# install aws-cli
RUN pip3 install awscli 

# install aws-sam-cli
RUN pip3 install aws-sam-cli==1.23.0

# change work directory
RUN mkdir -p /app
WORKDIR /app