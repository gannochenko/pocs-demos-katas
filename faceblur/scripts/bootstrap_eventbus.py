import os
import pika

def provision_rabbitmq_queue_and_exchange(host, exchange_name, queue_name, routing_key, exchange_type='direct'):
    try:
        # Connect to RabbitMQ server
        connection = pika.BlockingConnection(pika.ConnectionParameters(host=host))
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

provision_rabbitmq_queue_and_exchange(os.getenv('RABBITMQ_HOST'), os.getenv('RABBITMQ_EVENT_BUS_EXCHANGE_NAME'), os.getenv('RABBITMQ_EVENT_BUS_QUEUE_NAME'), os.getenv('RABBITMQ_EVENT_BUS_ROUTING_KEY'))
