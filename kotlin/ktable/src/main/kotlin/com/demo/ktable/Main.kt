package com.demo.ktable

import org.apache.kafka.streams.KafkaStreams
import org.apache.kafka.streams.StreamsBuilder
import org.apache.kafka.streams.StreamsConfig
import java.util.Properties

fun main() {
    println("Starting KTable Demo Application...")

    // Step 1: Configure Kafka Streams
    // These properties tell Kafka Streams where to connect and how to behave
    val props = Properties().apply {
        // Unique identifier for this application
        // Kafka uses this to manage state and consumer groups
        put(StreamsConfig.APPLICATION_ID_CONFIG, "ktable-demo-app")

        // Where is Kafka running?
        put(StreamsConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:9092")

        // How to serialize/deserialize keys and values
        // For now, we'll use simple strings
        put(StreamsConfig.DEFAULT_KEY_SERDE_CLASS_CONFIG, "org.apache.kafka.common.serialization.Serdes\$StringSerde")
        put(StreamsConfig.DEFAULT_VALUE_SERDE_CLASS_CONFIG, "org.apache.kafka.common.serialization.Serdes\$StringSerde")
    }

    // Step 2: Build the Topology
    // A topology is a graph of stream processors
    // Right now, it's empty - we'll add processing logic in next steps
    val builder = StreamsBuilder()

    // TODO: We'll add stream processing here in the next step
    println("Topology is empty for now - we'll add processing logic next")

    // Step 3: Create the KafkaStreams instance
    // This is the runtime that will execute our topology
    val topology = builder.build()
    val streams = KafkaStreams(topology, props)

    // Step 4: Add shutdown hook
    // Ensures graceful shutdown when you press Ctrl+C
    Runtime.getRuntime().addShutdownHook(Thread {
        println("\nShutting down gracefully...")
        streams.close()
    })

    // Step 5: Start the application
    println("Starting Kafka Streams...")
    println("Press Ctrl+C to stop")
    streams.start()
}
