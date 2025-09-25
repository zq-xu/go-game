from PIL import Image

# 打开原图并确保 RGBA
img = Image.open("wall.png").convert("RGBA")
w, h = img.size  # 32 × 32

# 裁剪底部 16 高的区域
bottom = img.crop((0, h//2, w, h))  # (left, top, right, bottom)

# 顺时针旋转 90 度
rotated = bottom.rotate(-90, expand=True)  # 顺时针旋转

# 新建画布（32×32，透明背景）
canvas = Image.new("RGBA", (w, h), (0, 0, 0, 0))

# 先贴原图
canvas.paste(img, (0, 0), mask=img)

# 再贴旋转后的底部
# 这里可以选择贴到某个位置，例如左上角 (0,0)：
canvas.paste(rotated, (0, 0), mask=rotated)

# 保存
canvas.save("wall_rotated.png", "PNG")

print("处理完成，生成 wall_rotated.png")
