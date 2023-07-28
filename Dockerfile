FROM node:16 as build-stage
WORKDIR /usr/src/app
COPY package*.json yarn.lock ./
RUN [ "/bin/bash", "-c", "yarn install --pure-lockfile 2> >(grep -v warning 1>&2) && mv node_modules ../"]
ENV PATH /usr/node_modules/.bin:$PATH
COPY . .
RUN yarn build

FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /usr/src/app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
