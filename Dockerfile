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

FROM golang:1.21.1 as appBuilder
WORKDIR /app/src/gosi/
COPY . .
WORKDIR /app/src/gosi/
RUN go build --trimpath -o app main.go


#CREATE FINAL IMAGE
############################################################################

FROM alpine:3.14 AS server

WORKDIR /app/josi

#COPY PLUGINS
# COPY --from=pluginBuilder /app/src/josi/mod/rpc/rpc.so /app/josi/mod/rpc/rpc.so

#COPY APP
COPY --from=appBuilder /app/src/gosi/app /app/gosi/app
COPY --from=appBuilder /app/src/gosi/gosi.db /app/gosi/gosi.db

WORKDIR /app/gosi
CMD ["./app"]

#EXPOSE 8080

#CMD ["gosandbox"]
