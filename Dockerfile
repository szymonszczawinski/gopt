# BUILD PLUGINS
############################################################################
#FROM golang:1.20.4 as pluginBuilder
#WORKDIR /app/src/josi/
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY . .
#WORKDIR /app/src/josi/mod
#RUN go build --buildmode=plugin --trimpath -o rpc/rpc.so rpc/rpc.go


#BUILD APP
############################################################################

FROM golang:1.21.3 as appBuilder
WORKDIR /app/src/gopt/
COPY . .
WORKDIR /app/src/gopt/
RUN go build --trimpath -o app main.go


#CREATE FINAL IMAGE
############################################################################

FROM alpine:3.14 AS server

WORKDIR /app/josi

#COPY PLUGINS
# COPY --from=pluginBuilder /app/src/josi/mod/rpc/rpc.so /app/josi/mod/rpc/rpc.so

#COPY APP
COPY --from=appBuilder /app/src/gopt/app /app/gopt/app
#COPY --from=appBuilder /app/src/gopt/gopt.db /app/gopt/gopt.db

WORKDIR /app/gopt
CMD ["./app"]

#EXPOSE 8080

#CMD ["gosandbox"]
