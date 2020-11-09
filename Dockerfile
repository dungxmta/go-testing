#
# Run test
#
FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir /opt/go
WORKDIR /opt/go
RUN mkdir go-testing

# install deps
WORKDIR /opt/go/go-testing
COPY go.mod .
RUN go mod download

# src code
COPY . .

RUN mkdir -p cov

# ; exit 0 to make sure docker not RUN fail when test FAIL
# otherwise docker will exit right away
RUN go test ./... -coverprofile=cov/cp.out > cov/cp.log; exit 0
#RUN go test ./... -coverprofile=cov/cp.out > cov/cp.log
RUN go tool cover -func=cov/cp.out -o cov/coverage.out

#
# Export output
#
FROM scratch as exporter
COPY --from=builder /opt/go/go-testing/cov /cov
# copy              <data_from_builder> to <output_in_exporter>

#
# demo
#

# DOCKER_BUILDKIT=1 docker build -f Dockerfile -o . .
#                                              -o <output_in_host> <cur_dir>
# copy <output_in_exporter> to <output_in_host>