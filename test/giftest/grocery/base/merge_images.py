from PIL import Image

def merge_images_with_spacing(image_paths, output_path, spacing=30, direction="horizontal"):
    """
    将多张图片按顺序拼接，间距可调，透明填充
    :param image_paths: 图片路径列表
    :param output_path: 输出路径
    :param spacing: 图片间距（像素）
    :param direction: 拼接方向 ("horizontal" 或 "vertical")
    """
    images = [Image.open(p).convert("RGBA") for p in image_paths]
    
    # 计算总宽高
    if direction == "horizontal":
        total_width = sum(img.width for img in images) + spacing * (len(images) - 1)
        max_height = max(img.height for img in images)
        result = Image.new("RGBA", (total_width, max_height), (0, 0, 0, 0))
        
        # 按顺序拼接
        x_offset = 0
        for img in images:
            result.paste(img, (x_offset, (max_height - img.height) // 2), mask=img)
            x_offset += img.width + spacing
    else:
        total_height = sum(img.height for img in images) + spacing * (len(images) - 1)
        max_width = max(img.width for img in images)
        result = Image.new("RGBA", (max_width, total_height), (0, 0, 0, 0))
        
        # 按顺序拼接
        y_offset = 0
        for img in images:
            result.paste(img, ((max_width - img.width) // 2, y_offset), mask=img)
            y_offset += img.height + spacing

    result.save(output_path, "PNG")
    print(f"✅ 拼接完成，保存为: {output_path}")


if __name__ == "__main__":
    image_files = ["1.png", "2.png", "3.png", "4.png"]  # 你的四张图片
    output_file = "merged.png"
    merge_images_with_spacing(image_files, output_file, spacing=30, direction="horizontal")
