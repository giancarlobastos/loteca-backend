FROM ubuntu:latest
ADD loteca-backend_unix .
CMD ["./loteca-backend_unix"]
EXPOSE 8080