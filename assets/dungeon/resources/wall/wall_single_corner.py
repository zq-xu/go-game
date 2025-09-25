from PIL import Image

# 打开原图并确保 RGBA
img = Image.open("wall.png").convert("RGBA")
w, h = img.size  # 32 × 32

# 右下角有效区域坐标
valid_left = w - 16   # 16
valid_top = h - 16    # 16
valid_right = w       # 32
valid_bottom = h      # 32

# 创建透明画布
canvas = Image.new("RGBA", (w, h), (0, 0, 0, 0))

# 裁剪有效区域
valid_region = img.crop((valid_left, valid_top, valid_right, valid_bottom))

# 把有效区域贴回右下角
canvas.paste(valid_region, (valid_left, valid_top), mask=valid_region)

# 保存
canvas.save("wall_single_corner.png", "PNG")

print("处理完成，生成 wall_single_corner.png")
