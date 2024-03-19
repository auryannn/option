FROM scratch
COPY go-option /
ENTRYPOINT ["/go-option"]
