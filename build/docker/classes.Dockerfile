FROM golang:1.19 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o out/ospc cmd/ospc/main.go
RUN go build -o out/classes cmd/classes/main.go

FROM alpine:3.16 AS run

COPY --from=build /app/out/ospc /usr/bin/ospc
COPY --from=build /app/out/classes /usr/bin/classes

ENTRYPOINT [ "/usr/bin/classes" ]
EXPOSE 8001
