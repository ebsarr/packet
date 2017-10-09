FROM alpine:latest as build

RUN apk -U --no-cache add \
	  build-base \
	  go \
	  git \
 	&& go get --ldflags '-linkmode external -extldflags "-static"' -u github.com/ebsarr/packet

FROM scratch
COPY --from=build /etc/ssl /etc/ssl
COPY --from=build /root/go/bin/packet /packet
ENTRYPOINT ["/packet"]
CMD ["-h"]
