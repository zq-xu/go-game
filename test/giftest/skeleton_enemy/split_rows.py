from PIL import Image

def split_image_rows(image_path, n, output_prefix="row"):
    """
    将图片均分为 n 行并保存
    :param image_path: 输入图片路径
    :param n: 行数
    :param output_prefix: 输出文件前缀
    """
    # 打开图片
    img = Image.open(image_path)
    width, height = img.size

    # 每行的高度
    row_height = height // n

    for i in range(n):
        # 计算裁剪区域
        top = i * row_height
        # 最后一行要保证包含剩余像素
        bottom = (i + 1) * row_height if i < n - 1 else height
        box = (0, top, width, bottom)
        row_img = img.crop(box)
        # 保存每行图片
        row_img.save(f"{output_prefix}_{i + 1}.png")
        print(f"保存: {output_prefix}_{i + 1}.png")

if __name__ == "__main__":
    image_path = "skeleton_enemy.png"  # 替换为你的图片路径
    n = 5  # 替换为你想分成的行数
    split_image_rows(image_path, n)
