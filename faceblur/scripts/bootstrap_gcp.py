import os
from google.cloud import storage


def create_bucket(bucket_name: str) -> None:
    client = storage.Client()
    bucket = client.bucket(bucket_name)
    bucket.location = 'eu'
    bucket.create()
    print(f"Bucket created: {bucket.name}")


create_bucket(os.getenv('STORAGE_IMAGE_BUCKET_NAME'))
