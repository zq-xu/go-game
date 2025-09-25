from PIL import Image

# 创建 32x32 的黑色图像（RGBA 模式）
black_img = Image.new("RGBA", (32, 32), (0, 0, 0, 255))  # (R,G,B,A)，255 表示不透明

# 保存为 PNG
black_img.save("black_tile.png", "PNG")

print("生成完成：black_tile.png")
