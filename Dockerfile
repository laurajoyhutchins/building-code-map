# Public BCM API image. Only compiled public software; private source data and
# operator-reviewed policy remain external at deployment time.
FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY provenance/ provenance/
COPY engine/ engine/
COPY httpapi/ httpapi/
COPY geocoder/ geocoder/
COPY utilitycontext/ utilitycontext/
COPY cmd/bcm/ cmd/bcm/
RUN CGO_ENABLED=0 go build -trimpath -o /bcm ./cmd/bcm

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /bcm /bcm
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/bcm"]
