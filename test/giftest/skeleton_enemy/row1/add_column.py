import cv2
import numpy as np

def insert_transparent_columns(image_path, insert_index, num_cols=20, output_path="output_with_gap.png"):
    """
    在指定列位置插入透明列
    :param image_path: 原图路径
    :param insert_index: 要插入的列索引（0 ~ 宽度-1）
    :param num_cols: 要插入的透明列数量
    :param output_path: 输出路径
    """
    # 读取图片，保持 alpha 通道
    img = cv2.imread(image_path, cv2.IMREAD_UNCHANGED)
    if img is None:
        print(f"无法读取图片: {image_path}")
        return

    h, w = img.shape[:2]

    # 如果图片没有 alpha 通道，添加 alpha
    if img.shape[2] == 3:
        alpha_channel = np.full((h, w), 255, dtype=np.uint8)
        img = np.dstack((img, alpha_channel))

    # 分割左右部分
    left_part = img[:, :insert_index]
    right_part = img[:, insert_index:]

    # 创建透明列
    transparent_cols = np.zeros((h, num_cols, 4), dtype=np.uint8)

    # 拼接
    new_img = np.hstack((left_part, transparent_cols, right_part))

    # 保存 PNG
    cv2.imwrite(output_path, new_img)
    print(f"保存完成: {output_path}, 新宽度: {new_img.shape[1]}")

if __name__ == "__main__":
    # 示例：在第100列插入20列透明列
    insert_transparent_columns("row_1.png", insert_index=577, num_cols=20)
