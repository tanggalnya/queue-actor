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
2. Start app
   ```shell
   go run main.app
   ```
3. Start worker
   ```shell
   go run worker
   ```