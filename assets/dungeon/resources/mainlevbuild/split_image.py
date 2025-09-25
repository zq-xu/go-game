import cv2
import os

# 输入输出路径
input_path = "mainlevbuild.png"
output_dir = "output_slices"

# 确保输出目录存在
os.makedirs(output_dir, exist_ok=True)

# 读取图片，带透明通道
img = cv2.imread(input_path, cv2.IMREAD_UNCHANGED)

# 如果有透明通道，取 alpha 通道做掩码
if img.shape[2] == 4:
    alpha = img[:, :, 3]
    gray = alpha
else:
    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

# 二值化
_, thresh = cv2.threshold(gray, 1, 255, cv2.THRESH_BINARY)

# 找到所有轮廓
contours, _ = cv2.findContours(thresh, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)

for i, cnt in enumerate(contours):
    x, y, w, h = cv2.boundingRect(cnt)
    # 裁剪出元素
    element = img[y:y+h, x:x+w]

    # 保存为 png
    output_path = os.path.join(output_dir, f"element_{i+1}.png")
    cv2.imwrite(output_path, element)

print(f"切割完成，共分成 {len(contours)} 个小图，保存在 {output_dir}/")
