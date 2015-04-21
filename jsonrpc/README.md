If I'm right, In `go net/rpc` `request.Seq` and `resposne.Seq` are exist for a synchronization between response and reply.
Like 'websocket' some transports should not guarantee about whether the request is belong to this reponse or not.
