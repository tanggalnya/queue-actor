# queue-actor

## Setup
### Rabbit MQ
We use https://github.com/wagslane/go-rabbitmq
1. Start
    ```bash
    docker run -d --name rabbitmq --hostname my-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
    ```