FROM golang:latest as build-env
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go build -o main .

FROM gcr.io/distroless/base
COPY --from=build-env /app/main /main
EXPOSE 8383
CMD ["/main"]