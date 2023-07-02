# Build stage
FROM node:lts as build-stage

WORKDIR /app

COPY package*.json ./
RUN yarn

COPY . .
RUN yarn build

# Run stage
FROM node:lts as production-stage
WORKDIR /app

COPY --from=build-stage /app .
CMD ["node", "./.output/server/index.mjs"]
