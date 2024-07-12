FROM node:18-alpine as builder

WORKDIR /product-service

COPY ./product-service/package*.json .

COPY ./product-service/ .

COPY ./proto ./proto

RUN npm install
# RUN npm install -g ts-proto
RUN npm install -g @nestjs/cli
RUN npm i --save @nestjs/microservices
RUN npm i --save @grpc/grpc-js @grpc/proto-loader

RUN npm run build

RUN npm prune --production

EXPOSE 50052

CMD [ "nest", "start" ]

# FROM node:18-alpine

# WORKDIR /product-service

# COPY --from=builder /product-service/node_modules ./node_modules
# COPY --from=builder /product-service/dist ./dist
# COPY --from=builder /product-service/package*.json ./

# COPY --from=builder /proto ./

# EXPOSE 50052

# CMD [ "node", "dist/main.js" ]