FROM alpine:3.4
RUN apk -U add ca-certificates
EXPOSE 9090
ADD bin/registry /bin/registry
CMD ["bin/registry", "serve", "--grpcPort", "9090"]