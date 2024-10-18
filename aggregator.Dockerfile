# Use a  golang alpine as the base image
FROM public.ecr.aws/docker/library/golang:1.23.0-alpine3.19 AS go_builder

RUN apk update
RUN apk add make cmake git alpine-sdk linux-headers

# Setup
# Read arguments
ARG ARCH=x86_64
ARG BUILD_DATE
ARG GIT_SHA
ARG VERSION

# Set env variables
ENV arch=$ARCH
ENV build_date=$BUILD_DATE
ENV commit_hash=$GIT_SHA
ENV service_name=aggregator
ENV version=$VERSION
RUN echo "building service: ${service_name}, version: ${version}, build date: ${build_date}, commit hash: ${commit_hash}, architecture: ${arch}"

# Set the working directory
COPY . /zenrock-avs
WORKDIR /zenrock-avs/aggregator/cmd

# Download dependencies
RUN go mod download
RUN go build -o aggregator

############################################################################################################

#SSL certs
FROM alpine:3.18.0 AS certs
RUN apk add --no-cache ca-certificates
RUN adduser -Ds /bin/bash appuser

# Copy binary to a scratch container. Let's keep our images nice and small!
COPY --from=go_builder /zenrock-avs/aggregator/cmd/aggregator /aggregator

# Set user
USER appuser

# Run the binary
ENTRYPOINT [ "/aggregator" ]
