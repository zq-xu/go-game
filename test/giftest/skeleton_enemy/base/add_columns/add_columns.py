import cv2
import numpy as np
import os

def insert_transparent_columns_multi(image_path, insert_indices, num_cols=20):
    """
    在指定多个列位置插入透明列，并自动生成输出文件名
    :param image_path: 原图路径
    :param insert_indices: 要插入的列索引列表，例如 [100, 300, 500]
    :param num_cols: 每个位置插入的透明列数量
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

    # 按插入位置排序
    insert_indices = sorted(insert_indices)

    offset = 0  # 每次插入后的索引偏移
    for index in insert_indices:
        insert_index = index + offset

        # 分割左右部分
        left_part = img[:, :insert_index]
        right_part = img[:, insert_index:]

        # 创建透明列
        transparent_cols = np.zeros((h, num_cols, 4), dtype=np.uint8)

        # 拼接
        img = np.hstack((left_part, transparent_cols, right_part))

        offset += num_cols  # 后续插入需要加上已经插入的列数

    # 生成输出路径：原名 + _add_columns + 原扩展名
    base, ext = os.path.splitext(image_path)
    output_path = f"{base}_add_columns{ext}"

    # 保存图片
    cv2.imwrite(output_path, img)
    print(f"保存完成: {output_path}, 新宽度: {img.shape[1]}")

if __name__ == "__main__":
    insert_transparent_columns_multi("row_1.png", insert_indices=[327,577], num_cols=20)