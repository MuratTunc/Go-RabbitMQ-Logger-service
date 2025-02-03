

# Docker-compose file
docker-compose.yml makes all of this work by ensuring that:

    RabbitMQ runs on a separate container, configured to accept connections.
    MongoDB runs on its own container, storing data for your Logger Service.
    Logger Service (your Go app) depends on both RabbitMQ and MongoDB and communicates with them via their respective service names (mongo and rabbitmq).

# Architecture in Docker   
RabbitMQ will be responsible for message queues. Your Logger Service will send logs to RabbitMQ.
MongoDB will store your logs. Your Logger Service will take messages from RabbitMQ and store them in MongoDB.