#!/bin/sh
sudo echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
sudo apt-get update
sudo apt-get install google-cloud-sdk -y
brew install --build-from-source yq
helm plugin install https://github.com/chartmuseum/helm-push.git
helm repo add ysma500 http://ysma500.tk:314
helm repo update
