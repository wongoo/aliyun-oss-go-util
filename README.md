# Usage

```
go get github.com/wongoo/aliyun-oss-go-util

# -----
# delete objects with access urls containing in a file
aliyun-oss-go-util -access_key_id=<ACCESS_KEY_ID> -access_key_secret=<ACCESS_KEY_SECRET> \
	-endpoint=http://oss-cn-shenzhen.aliyuncs.com \
	-bucket=<MY_BUCKET_NAME> \
	-action=batch_delete_url_in_file \
	-file_path=~/data_with_file_urls.txt \
	-url_prefix=http://<MY_BUCKET_NAME>.oss-cn-shenzhen.aliyuncs.com


```