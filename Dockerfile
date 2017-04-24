FROM golang:1.8.1
COPY bin .
EXPOSE 8081
CMD ["bash", "linux"]
