cd /root/k8s/web/back/k8s/
docker image rm k8s_demo/back:1.0
cd ../
docker build -t k8s_demo/back:1.0 .
cd ./k8s/
docker save  k8s_demo/back -o k8s_demo_back.tar
ctr -n k8s.io images import k8s_demo_back.tar