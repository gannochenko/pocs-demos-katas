#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import requests
import threading
import time
import random


def main():
    # Define the URLs
    url1 = "https://www.example1.com"
    url2 = "https://www.example2.com"

    # Create threads
    thread1 = threading.Thread(target=execute_request, args=(url1,))
    thread2 = threading.Thread(target=execute_request, args=(url2,))

    # Start the threads
    thread1.start()
    thread2.start()


# Define the function for the HTTP calls
def execute_request(url):
    while True:
        response = requests.get(url)
        print(f"Response from {url}: {response.status_code}")
        pause_time = random.randint(1, 10)  # Random pause between 1 to 10 seconds
        time.sleep(pause_time)


if __name__ == "__main__":
    main()
