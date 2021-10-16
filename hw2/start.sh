#build image
docker build -t wangciao/geektime_cloud_hw_hw2 .

#start server
docker run -d -p 8080:80 wangciao/geektime_cloud_hw_hw2
