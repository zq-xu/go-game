import cv2
import numpy as np
import os

def flip_columns_auto(image_path, min_gap=10):
    """
    自动识别人物列，并在每列内部水平翻转
    输出文件名 = 原名 + _horizontal_flipped + 原扩展名
    :param image_path: 输入图片（支持透明）
    :param min_gap: 人物列之间最少透明像素数，用于分割列
    """
    img = cv2.imread(image_path, cv2.IMREAD_UNCHANGED)
    if img is None:
        print(f"无法读取图片: {image_path}")
        return

    h, w = img.shape[:2]

    # 确保有 alpha 通道
    if img.shape[2] == 3:
        alpha = np.full((h, w), 255, dtype=np.uint8)
        img = np.dstack((img, alpha))

    # 获取 alpha 通道
    alpha = img[:, :, 3]

    # 计算每列非透明像素数
    col_sum = np.sum(alpha > 0, axis=0)

    # 找连续非零列段
    col_boundaries = []
    start = None
    for i, val in enumerate(col_sum):
        if val > 0:
            if start is None:
                start = i
        else:
            if start is not None:
                end = i - 1
                if end - start + 1 >= 1:
                    col_boundaries.append([start, end])
                start = None
    if start is not None:
        col_boundaries.append([start, w-1])

    # 合并间隔小于 min_gap 的列段
    merged_boundaries = []
    prev_start, prev_end = col_boundaries[0]
    for s, e in col_boundaries[1:]:
        if s - prev_end <= min_gap:
            prev_end = e
        else:
            merged_boundaries.append([prev_start, prev_end])
            prev_start, prev_end = s, e
    merged_boundaries.append([prev_start, prev_end])

    # 对每列内部水平翻转
    result_img = img.copy()
    for start, end in merged_boundaries:
        col = img[:, start:end+1]
        col_flipped = cv2.flip(col, 1)
        result_img[:, start:end+1] = col_flipped

    # 生成输出文件名：原名 + _horizontal_flipped + 原扩展名
    base, ext = os.path.splitext(image_path)
    output_path = f"{base}_horizontal_flipped{ext}"

    cv2.imwrite(output_path, result_img)
    print(f"保存完成: {output_path}")

if __name__ == "__main__":
    # 自动识别人物列并翻转
    flip_columns_auto("row_1_add_columns.png", min_gap=10)
