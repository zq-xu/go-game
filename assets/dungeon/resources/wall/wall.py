from PIL import Image

# 打开原图并确保是 RGBA 模式
img = Image.open("element_58.png").convert("RGBA")
w, h = img.size  # 原图大小，比如 64x16

new_h = 32
new_w = w

# 创建完全透明的画布（RGBA 全为 0）
canvas = Image.new("RGBA", (new_w, new_h), (0, 0, 0, 0))

# 先清理画布（这一步其实不是必须的，因为 new 已经是全 0，但加一层保险）
pixels = canvas.load()
for y in range(new_h - h):
    for x in range(new_w):
        pixels[x, y] = (0, 0, 0, 0)  # 确保顶部区域没有“幽灵像素”

# 把原图贴到底部，使用 alpha 通道作为 mask
canvas.paste(img, (0, new_h - h), mask=img)

# 裁掉左半部分（保留右半边）
final = canvas.crop((0, 0, new_w // 2, new_h))

# 保存为 PNG，必须使用 RGBA 才能保留透明度
final.save("wall.png", "PNG")
