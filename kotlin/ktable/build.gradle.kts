plugins {
    kotlin("jvm") version "2.1.0"
    application
}

group = "com.demo"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    // Kafka Streams
    implementation("org.apache.kafka:kafka-streams:3.6.1")
    implementation("org.apache.kafka:kafka-clients:3.6.1")

    // Kotlin standard library
    implementation("org.jetbrains.kotlin:kotlin-stdlib")

    // Logging
    implementation("ch.qos.logback:logback-classic:1.4.14")
    implementation("org.slf4j:slf4j-api:2.0.9")
}

// Configure JVM toolchain - Gradle will auto-download JDK 21 if needed
kotlin {
    jvmToolchain(21)
}

application {
    mainClass.set("com.demo.ktable.MainKt")
}
