FROM golang:1.20-alpine as build
WORKDIR /work
COPY . .
RUN go build -o app

FROM alpine as runtime
COPY --from=build /work/app .

CMD ./app
