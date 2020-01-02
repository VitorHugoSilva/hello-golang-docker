FROM scratch

COPY bin/projeto_go /projeto_go

ENTRYPOINT ["/projeto_go"]