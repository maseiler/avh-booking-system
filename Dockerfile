# build client spa with npm
FROM node:lts-alpine as build-npm
WORKDIR /app
COPY ./client .
RUN npm install
RUN npm run build

# build server binary
FROM golang:alpine as build-go
RUN apk --no-cache add git build-base 
WORKDIR $GOPATH/src/github.com/maseiler/avh-booking-system/server
COPY server ./
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux
RUN echo "Go path: $GOPATH"
RUN env GOOS=linux GOARCH=386 go build -o /avh_bs_main main.go
RUN chmod +x /avh_bs_main

# the actual app stage
FROM alpine
WORKDIR /app
COPY --from=build-npm /app/dist /app/dist
COPY --from=build-go /avh_bs_main /app/avh_bs_main
EXPOSE 8081

# default env variables
ENV AVHBS_DB_USER db_user
ENV AVHBS_DB_PASS db_pass
ENV AVHBS_DB_NAME avhbs_db
ENV AVHBS_DB_IP db
ENV AVHBS_DB_PORT 3306
ENV AVHBS_FRONTEND_PATH /app/dist/

RUN echo -e '#!/bin/sh\necho "Sleeping..."\nsleep 45\necho "Starting app"!\n/app/avh_bs_main\n' > /start.sh && chmod +x /start.sh && cat /start.sh

#execute server
CMD ["sh", "/start.sh"]
