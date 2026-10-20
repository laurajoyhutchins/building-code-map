FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY engine/ engine/
COPY httpapi/ httpapi/
COPY geocoder/ geocoder/
COPY cmd/bcm/ cmd/bcm/
RUN CGO_ENABLED=0 go build -trimpath -o /bcm ./cmd/bcm

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /bcm /bcm
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/bcm"]
