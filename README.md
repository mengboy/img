### jpeg、jpg、png批量压缩转换工具

#### 图片格式转换
png和jpg/jpeg互转
```shell
# 将image下的png图片批量转换为jpeg图片并压缩质量到80
./img convert -s /images -d /imagesjpeg -c jpeg -q 80

# 将image下的test.png转换为test.jpeg并压缩质量到80 
./img convert -s /images/test.png -d /images/test.jpeg -c jpeg -q 80
```
#### 图片尺寸缩小
jpeg、jpg、png 图片尺寸缩小
```shell
# 将image下的png图片尺寸批量缩小至原来的80%
./img resize -s /images -d /images8 -f 0.8

# 将image下的test.png图片尺寸缩小至原来的80% 
./img resize -s /images/test.png -d /images/test8.png -f 0.8
```

#### 图片质量压缩
jpeg、jpg、png 图片压缩
```shell
# 将image下的图片批量压缩质量到80
./img resize -s /images -d /images8 -q 80

# 将image下的test.jpeg压缩质量到80 
./img resize -s /images/test.jpeg -d /images/test8.jpeg - 80
```

只有jpg\jpeg支持压缩，png为无损图片格式只支持Go默认的压缩压缩等级
```go
encoder := png.Encoder{
		CompressionLevel: png.BestCompression,
	}
```