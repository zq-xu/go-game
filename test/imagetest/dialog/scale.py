from PIL import Image

def resize_image(input_path, output_path, max_size):
    """
    等比例缩小图片，最长边不超过 max_size
    :param input_path: 输入图片路径
    :param output_path: 输出图片路径
    :param max_size: 缩小后最长边的像素大小（例如 800）
    """
    with Image.open(input_path) as img:
        # 获取原始尺寸
        width, height = img.size

        # 计算缩放比例
        ratio = min(max_size / width, max_size / height)

        # 如果图片本身已经比目标小，就不缩放
        if ratio >= 1:
            print("图片已经比目标尺寸小，无需缩放。")
            img.save(output_path)
            return

        # 计算新的尺寸
        new_size = (int(width * ratio), int(height * ratio))
        resized_img = img.resize(new_size, Image.Resampling.LANCZOS)

        # 保存输出文件
        resized_img.save(output_path)
        print(f"已保存到 {output_path}，新尺寸为 {new_size}")

# 示例用法
resize_image("origin.png", "dialog_icon.png", 20)
