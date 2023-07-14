<h1 style="font-weight:normal">
  <a href="">
    <img src="https://cdn.jsdelivr.net/gh/NeeDKK/cloudimg@master/data/logo.png" alt="needkk" width=35>
  </a>
  &nbsp;tes-face
  <a href="http://face.needkk.com/"><img src="https://img.shields.io/static/v1?label=tes-face&color=green&message=start now"/></a>
  <a href="http://face.needkk.com/" target="_blank"><img src="https://img.shields.io/static/v1?label=release&color=blue&message=latest"/></a>

</h1>
Demo: <a href="http://face.needkk.com/" target="_blank">http://face.needkk.com/</a>
<br/>
(单核服务器，上传视频资源时请选择小于10M或小于20s的视频文件，否则会超时！！！)
<hr>
通过Ubuntu作为基础镜像构建拥有opencv+golang+dlib环境，通过golang调用opencv和dlib进行人脸识别。
golang基础包 go-face + gocv-go 通过两个开源包分别处理人脸识别图片打码和视频打码。<br/>
基础构件仓库地址： <a href="https://github.com/NeeDKK/face-open-base" target="_blank">https://github.com/NeeDKK/face-open-base</a>

## 参考仓库
>https://github.com/Kagami/go-face (go-face仓库)
> 
> https://github.com/Kagami/go-face-testdata （图片模型仓库）
> 
> https://gocv.io/ （gocv-go官方地址）
> 
> https://github.com/opencv/opencv/tree/master/data/haarcascades （opencv模型仓库）



## 构建自己的opencv+go+dlib golang后端应用镜像

>
> ```shell
> sh build-tes-face.sh
> ```
>

## 构建自己的前端镜像

>
> ```shell
> cd test-face-web
> docker login -u xxxx -p xxxx
> docker build -f Dockerfile -t tes-face-web:latest .
> docker tag tes-face-web:latest xxxx/tes-face-web:latest
> docker push xxxx/tes-face-web:latest
> ```
>

## TODO :

> 1.优化人脸识别单图片识别的模型，提高识别率（目前脸部有遮挡或全侧脸有概率无法识别）
> 
> 2.优化视频打码的运行效率，通过协程多帧视频并发处理