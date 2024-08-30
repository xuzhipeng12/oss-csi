#!/bin/bash

echo "[1/3] build"
# pull and build
cd /data/csi/oss-csi
git pull  &&  go build -o bin/
ls -lah   bin/

echo "[2/3] docker image build"
# docker image  build & push
cd /data/csi/oss-csi/bin/
docker build -t xuzhipeng12/xzp-oss:0.1  . && docker push xuzhipeng12/xzp-oss:0.1

echo "[3/3] restart service"
# restart csi pods
kubectl get  pod  | grep oss-csi- | awk '{system("kubectl delete pod "$1)}'

echo "all done! "
