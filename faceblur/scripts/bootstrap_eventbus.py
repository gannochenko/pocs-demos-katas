import os
import pika
import urllib.parse

def parse_rabbitmq_dsn(dsn):
    """Parses a RabbitMQ DSN and returns connection parameters."""
    parsed_url = urllib.parse.urlparse(dsn)
    
    return {
        "host": parsed_url.hostname or "localhost",
        "port": parsed_url.port or 5672,
        "username": parsed_url.username or "guest",
        "password": parsed_url.password or "guest",
        "vhost": parsed_url.path.lstrip('/') or "/"
    }

def provision_rabbitmq_queue_and_exchange(dsn, exchange_name, queue_name, routing_key, exchange_type='direct'):
    try:
        # Parse the DSN
        params = parse_rabbitmq_dsn(dsn)

        # Connect to RabbitMQ server using DSN parameters
        credentials = pika.PlainCredentials(params["username"], params["password"])
        connection_params = pika.ConnectionParameters(
            host=params["host"],
            port=params["port"],
            virtual_host=params["vhost"],
            credentials=credentials
        )

        connection = pika.BlockingConnection(connection_params)
        channel = connection.channel()

        # Declare an exchange
        channel.exchange_declare(exchange=exchange_name, exchange_type=exchange_type, durable=True)

        # Declare a queue
        channel.queue_declare(queue=queue_name, durable=True)

        # Bind the queue to the exchange with a routing key
        channel.queue_bind(exchange=exchange_name, queue=queue_name, routing_key=routing_key)

        print(f"Successfully provisioned exchange '{exchange_name}' and queue '{queue_name}' with routing key '{routing_key}'.")

        # Close the connection
        connection.close()
    except Exception as e:
        print(f"Failed to provision RabbitMQ components: {e}")

# Use RabbitMQ DSN from the environment variable
RABBITMQ_DSN = os.getenv('RABBITMQ_DSN', 'amqp://guest:guest@localhost:5672/')

provision_rabbitmq_queue_and_exchange(
    RABBITMQ_DSN,
    os.getenv('RABBITMQ_EVENT_BUS_EXCHANGE_NAME'),
    os.getenv('RABBITMQ_EVENT_BUS_QUEUE_NAME'),
    os.getenv('RABBITMQ_EVENT_BUS_ROUTING_KEY')
)
