FROM alpine:3.21.3
RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Taipei
WORKDIR /app
COPY .env stress-cpu /app/
CMD [ "./stress-cpu" ]