import onnx

model_path = "/Users/gannochenko/.cache/huggingface/hub/models--AdamCodd--YOLOv11n-face-detection/snapshots/895c5d8452685fab8ff9e33b28098d926adc90c4/model.onnx"

model = onnx.load(model_path)
for output in model.graph.output:
    print(f"Output Name: {output.name}, Shape: {output.type.tensor_type.shape}")

