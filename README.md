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
(���˷��������ϴ���Ƶ��Դʱ��ѡ��С��10M��С��20s����Ƶ�ļ�������ᳬʱ������)
<hr>
ͨ��Ubuntu��Ϊ�������񹹽�ӵ��opencv+golang+dlib������ͨ��golang����opencv��dlib��������ʶ��
golang������ go-face + gocv-go ͨ��������Դ���ֱ�������ʶ��ͼƬ�������Ƶ���롣<br/>
���������ֿ��ַ�� <a href="https://github.com/NeeDKK/face-open-base" target="_blank">https://github.com/NeeDKK/face-open-base</a>

## �ο��ֿ�
>https://github.com/Kagami/go-face (go-face�ֿ�)
> 
> https://github.com/Kagami/go-face-testdata ��ͼƬģ�Ͳֿ⣩
> 
> https://gocv.io/ ��gocv-go�ٷ���ַ��
> 
> https://github.com/opencv/opencv/tree/master/data/haarcascades ��opencvģ�Ͳֿ⣩



## �����Լ���opencv+go+dlib golang���Ӧ�þ���

>
> ```shell
> sh build-tes-face.sh
> ```
>

## �����Լ���ǰ�˾���

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

> 1.�Ż�����ʶ��ͼƬʶ���ģ�ͣ����ʶ���ʣ�Ŀǰ�������ڵ���ȫ�����и����޷�ʶ��
> 
> 2.�Ż���Ƶ���������Ч�ʣ�ͨ��Э�̶�֡��Ƶ��������