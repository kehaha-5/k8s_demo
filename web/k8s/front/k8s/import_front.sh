cd /root/k8s/web/front/k8s/
docker image rm k8s_demo/front:1.0
cd ../
docker build -t k8s_demo/front:1.0 .
cd ./k8s/
docker save  k8s_demo/front -o k8s_demo_front.tar
ctr -n k8s.io images import k8s_demo_front.tar