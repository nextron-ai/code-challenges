import easyocr

reader = easyocr.Reader(["pt"])
result = reader.readtext("./img1.png")

for item in result:
    bounding_box = item[0]
    text = item[1]
    confidence_level = item[2]
