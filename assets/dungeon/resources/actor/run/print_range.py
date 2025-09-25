import cv2
import numpy as np

def find_bounding_boxes(image_path):
    # 读取图像
    img = cv2.imread(image_path)
    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

    # 二值化（你可能要调阈值，或者用自适应阈值/大津法）
    _, thresh = cv2.threshold(gray, 1, 255, cv2.THRESH_BINARY)

    # 找连通域
    num_labels, labels, stats, centroids = cv2.connectedComponentsWithStats(thresh, connectivity=8)

    bboxes = []
    for i in range(1, num_labels):  # 跳过背景 (label=0)
        x, y, w, h, area = stats[i]
        bboxes.append((x, y, x+w, y+h))

    return bboxes

if __name__ == "__main__":
    path = "run_right.png"
    boxes = find_bounding_boxes(path)
    for i, box in enumerate(boxes, start=1):
        print(f"人物 {i} 像素范围: (xmin={box[0]}, ymin={box[1]}, xmax={box[2]}, ymax={box[3]})")
