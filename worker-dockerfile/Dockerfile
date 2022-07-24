FROM golang:1.18 as builder

WORKDIR /worker

COPY . .

RUN make worker

#FROM alpine:latest
#COPY --from=builder /app/build/worker ./

CMD ["build/worker", "start"]