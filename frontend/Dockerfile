FROM node:latest

WORKDIR /app

RUN npm install -g pnpm

COPY ./package.json ./
COPY ./pnpm-lock.yaml ./

RUN pnpm install

COPY . .


RUN pnpm run build
CMD ["pnpm", "run", "start"]
