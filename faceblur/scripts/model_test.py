from huggingface_hub import hf_hub_download
from ultralytics import YOLO
from PIL import Image
from supervision import Detections

model_path = hf_hub_download(repo_id="AdamCodd/YOLOv11n-face-detection", filename="model.pt")
model = YOLO(model_path)

file = "/Users/gannochenko/proj/tryout/onnxruntime_go_examples/image_object_detect/input.png"

output = model(Image.open(file))
resultss = Detections.from_ultralytics(output[0])
print(resultss)

# results = model.predict(file, save=True) # saves the result in runs/detect/predict

# print(results)
