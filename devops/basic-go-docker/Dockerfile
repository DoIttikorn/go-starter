# เอา images 1.19 ให้ชื่อว่า build-base
FROM golang:1.19-alpine as build-base

# สร้าง directory ที่จะเก็บ source code
WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/go-simple .

# ==================================

FROM alpine:3.16.2
# copy data from old image to new image [old image path] [new image path]
COPY --from=build-base /app/out/go-simple app/go-simple

CMD ["/app/go-simple"]