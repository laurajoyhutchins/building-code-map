# Public BCM API image. Only compiled public software; private source data and
# operator-reviewed policy remain external at deployment time.
FROM golang:1.25-alpine@sha256:1ae0735f00daffa3aaf1363a5184c0d2dc55c78e3db4ec70241cdac97bf84b59 AS build
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

FROM gcr.io/distroless/static-debian12:nonroot@sha256:afa5c872c891853ca7fcf1f12c3edb23f7eeef36189728842dd51042ff57f7ab
COPY --from=build /bcm /bcm
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/bcm"]
