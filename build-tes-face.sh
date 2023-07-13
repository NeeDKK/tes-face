docker login -u needkk -p 110120119wzj
docker build -f DockerfileBuild -t tes-face:latest .
docker tag tes-face:latest needkk/tes-face:latest
docker push needkk/tes-face:latest