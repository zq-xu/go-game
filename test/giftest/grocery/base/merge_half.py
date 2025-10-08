from PIL import Image

def merge_images_with_spacing(image_paths, output_path, spacing=30, direction="horizontal", scale=0.7):
    """
    将多张图片按顺序拼接，间距可调，透明填充，并支持等比缩放
    :param image_paths: 图片路径列表
    :param output_path: 输出路径
    :param spacing: 图片间距（像素）
    :param direction: 拼接方向 ("horizontal" 或 "vertical")
    :param scale: 缩放比例（0.5 表示缩小一半）
    """
    images = []

    # 打开并缩放图片
    for p in image_paths:
        img = Image.open(p).convert("RGBA")
        if scale != 1.0:
            new_size = (int(img.width * scale), int(img.height * scale))
            img = img.resize(new_size, Image.Resampling.LANCZOS)
        images.append(img)

    # 计算总宽高
    if direction == "horizontal":
        total_width = sum(img.width for img in images) + spacing * (len(images) - 1)
        max_height = max(img.height for img in images)
        result = Image.new("RGBA", (total_width, max_height), (0, 0, 0, 0))

        # 按顺序拼接
        x_offset = 0
        for img in images:
            y_offset = (max_height - img.height) // 2  # 居中对齐
            result.paste(img, (x_offset, y_offset), mask=img)
            x_offset += img.width + spacing
    else:
        total_height = sum(img.height for img in images) + spacing * (len(images) - 1)
        max_width = max(img.width for img in images)
        result = Image.new("RGBA", (max_width, total_height), (0, 0, 0, 0))

        # 按顺序拼接
        y_offset = 0
        for img in images:
            x_offset = (max_width - img.width) // 2  # 居中对齐
            result.paste(img, (x_offset, y_offset), mask=img)
            y_offset += img.height + spacing

    result.save(output_path, "PNG")
    print(f"✅ 拼接完成，保存为: {output_path}")


if __name__ == "__main__":
    image_files = ["1.png", "2.png", "3.png", "4.png"]
    output_file = "merged_scaled.png"
    merge_images_with_spacing(image_files, output_file, spacing=30, direction="horizontal", scale=0.6)
