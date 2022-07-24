FROM golang:1.18 as builder

WORKDIR /server

COPY . .

RUN make server

#FROM alpine:latest
#COPY --from=builder /app/build/server ./

EXPOSE 3000
CMD ["build/server", "start"]