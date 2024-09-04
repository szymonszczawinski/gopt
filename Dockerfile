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

FROM golang:1.23.0-alpine3.20 as build-stage
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /gopt

#CREATE FINAL IMAGE
############################################################################

FROM alpine:3.19 AS release-stage

WORKDIR /

#COPY PLUGINS
# COPY --from=pluginBuilder /app/src/josi/mod/rpc/rpc.so /app/josi/mod/rpc/rpc.so

#COPY APP
COPY --from=build-stage /gopt /gopt
#COPY --from=appBuilder /app/src/gopt/gopt.db /app/gopt/gopt.db
EXPOSE 8081
ENV DB_URL=postgresql://postgres:postgres@172.18.0.1:5432/gopt
# WORKDIR /app/gopt
CMD "/gopt"


#CMD ["gosandbox"]
