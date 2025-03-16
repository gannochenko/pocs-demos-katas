package image

// FilterDistinctBoxes implements Non-Maximum Suppression (NMS) to filter
// bounding boxes that have significant overlap with others.
// Parameters:
//   - boxes: The array of bounding boxes to filter
//   - iouThreshold: Boxes with IoU above this threshold will be considered duplicates (typically 0.45-0.7)
//   - confidenceThreshold: Optional minimum confidence (boxes below this are filtered, use 0 to ignore)
//   - maxBoxes: Optional maximum number of boxes to return (use 0 for no limit)
//
// Returns a filtered array of the most distinct bounding boxes.
func FilterDistinctBoxes(boxes []BoundingBox, iouThreshold, confidenceThreshold float32, maxBoxes int) []BoundingBox {
    // If no boxes or just one box, return as is
    if len(boxes) <= 1 {
        return boxes
    }

    // Filter by confidence threshold
    var filteredBoxes []BoundingBox
    if confidenceThreshold > 0 {
        for _, box := range boxes {
            if box.Confidence >= confidenceThreshold {
                filteredBoxes = append(filteredBoxes, box)
            }
        }
    } else {
        filteredBoxes = make([]BoundingBox, len(boxes))
        copy(filteredBoxes, boxes)
    }

    // Sort boxes by confidence (highest first)
    // This ensures we keep boxes with higher confidence scores
    sortBoxesByConfidence(filteredBoxes)

    // Apply Non-Maximum Suppression
    var selectedBoxes []BoundingBox

    for len(filteredBoxes) > 0 {
        // Take the box with highest confidence
        currentBox := filteredBoxes[0]
        selectedBoxes = append(selectedBoxes, currentBox)
        
        // Check if we've reached the maximum number of boxes
        if maxBoxes > 0 && len(selectedBoxes) >= maxBoxes {
            break
        }

        // Remove the current box from the list
        filteredBoxes = filteredBoxes[1:]
        
        // Create a new list without boxes that have high IoU with the current box
        var remainingBoxes []BoundingBox
        for _, box := range filteredBoxes {
            // If boxes don't overlap too much, keep them
            boxPtr, currentBoxPtr := &box, &currentBox
            if boxPtr.iou(currentBoxPtr) <= iouThreshold {
                remainingBoxes = append(remainingBoxes, box)
            }
            // Otherwise, they're considered duplicates and are dropped
        }
        
        filteredBoxes = remainingBoxes
    }

    return selectedBoxes
}

// sortBoxesByConfidence sorts the boxes by confidence in descending order
func sortBoxesByConfidence(boxes []BoundingBox) {
    // Simple bubble sort (for small arrays this is fine)
    // In production, you might want to use sort.Slice instead
    for i := 0; i < len(boxes)-1; i++ {
        for j := i + 1; j < len(boxes); j++ {
            if boxes[i].Confidence < boxes[j].Confidence {
                boxes[i], boxes[j] = boxes[j], boxes[i]
            }
        }
    }
}

// FilterBoundingBoxesByClass applies Non-Maximum Suppression per class label
// This is useful when you want to keep the best boxes for each class separately
func FilterBoundingBoxesByClass(boxes []BoundingBox, iouThreshold, confidenceThreshold float32, maxBoxesPerClass int) []BoundingBox {
    // Group boxes by class label
    classBuckets := make(map[string][]BoundingBox)
    for _, box := range boxes {
        classBuckets[box.Label] = append(classBuckets[box.Label], box)
    }
    
    // Apply NMS to each class separately
    var result []BoundingBox
    for _, classBoxes := range classBuckets {
        filtered := FilterDistinctBoxes(classBoxes, iouThreshold, confidenceThreshold, maxBoxesPerClass)
        result = append(result, filtered...)
    }
    
    return result
}