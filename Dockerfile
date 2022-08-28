FROM golang:1.18
WORKDIR /builddir
ADD . .
RUN go build -o /app/my-go-backend
RUN mv /builddir/config/config.yaml /app/config/dev.yaml
RUN rm -r /builddir
ENV ENVIRONMENT_TYPE=DEV
EXPOSE 8043/tcp
ENTRYPOINT ["./app"]