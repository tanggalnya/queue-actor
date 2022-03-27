# queue-actor

## Setup
### Rabbit MQ
We use https://github.com/wagslane/go-rabbitmq

1. First time only
    ```sh
    docker run -d --name rabbitmq --hostname my-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
    ```
   Else execute
   ```sh
   docker start rabbitmq 
   ```
2. Build
   ```shell
   make build   
   ```
4. Start server
   ```shell
   ./out/queue-actor server
   ```
5. Start worker
   ```shell
   ./out/queue-actor worker
   ```