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
2. Dashboard
   1. http://localhost:15672/
   2. username:password: guest:guest
   
### App
1. Build
   ```shell
   make build   
   ```
2. Start server
   ```shell
   ./out/queue-actor server
   ```
3. Start worker
   ```shell
   ./out/queue-actor worker
   ```