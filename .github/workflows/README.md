# Supported Cloud for Deployment
In Ubivius deployment solution, we are currently supporting these cloud providers:
- Google Kubernetes Engine(GKE)
- Linode Kubernetes Engine(LKE)

To customize you can add the desired deployment code block at the end of your `push-{target_branch}.yml` GitHub Action file following the procedure described below.

## Config deployment steps
### Add GitHub Secrets
| Secret Name | Description | Example | Is it mandatory? |
| ----------- | ----------- | ------- | ---------------- |
| <span style="font-size:0.75em;">DOCKERHUB_USERNAME</span> | Your DockerHub service account short username used to push your built Docker container to DockerHub | <span style="font-size:0.75em;">your_name</span> | Yes |
| <span style="font-size:0.75em;">DOCKERHUB_PWD</span> | Your DockerHub service account's password to enable pushing built containers | <span style="font-size:0.75em;">abc123</span> | Yes |
| <span style="font-size:0.75em;">GKE_SERVICEACCOUNT_EMAIL</span> | This is your GKE service account's email | <span style="font-size:0.75em;">my_deployment_bot@cluster_name.iam.gserviceaccount.com</span> | Only if you want to deploy on **GKE** |
| <span style="font-size:0.75em;">GKE_SERVICEACCOUNT_JSONKEY</span> | The downloaded json file's content when you create your service account's key | <span style="font:normal 7pt Arial;">{<br/>  "type": "{service_account}",<br/>  "project_id": "{project_id}",<br/>    "private_key_id": "{key_id}",<br/>  "private_key": "-----BEGIN PRIVATE KEY-----\{your_key}\n-----END PRIVATE KEY-----\n",<br/>  "client_email": "{service_account}",<br/>  "client_id": "{user_id}",<br/>  "auth_uri": "https://accounts.google.com/o/oauth2/auth",<br/>  "token_uri": "https://oauth2.googleapis.com/token",<br/>  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",<br/>  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/{service_account}"<br/>}</span> | Only if you want to deploy on **GKE** |
| <span style="font-size:0.75em;">GKE_DEV_CLUSTER_NAME</span> | The name of your dev cluster | <span style="font-size:0.75em;">dev-environment</span> | Only if you want to deploy on **GKE** from dev branch to your dev cluster |
| <span style="font-size:0.75em;">GKE_DEV_CLUSTER_ZONE</span> | The zone your dev cluster is located | <span style="font-size:0.75em;">us-east1-c</span> | Only if you want to deploy on **GKE** from dev branch to your dev cluster |
| <span style="font-size:0.75em;">GKE_DEV_CLUSTER_PROJECT</span> | The name of the GKE project containing your dev cluster | <span style="font-size:0.75em;">your-project</span> | Only if you want to deploy on **GKE** from dev branch to your dev cluster |
| <span style="font-size:0.75em;">GKE_PROD_CLUSTER_NAME</span> | The name of your prod cluster | <span style="font-size:0.75em;">prod-environment</span> | Only if you want to deploy on **GKE** from main/master branch to your prod cluster |
| <span style="font-size:0.75em;">GKE_PROD_CLUSTER_ZONE</span> | The zone your prod cluster is located | <span style="font-size:0.75em;">us-east1-c</span> | Only if you want to deploy on **GKE** from main/master branch to your prod cluster |
| <span style="font-size:0.75em;">GKE_PROD_CLUSTER_PROJECT</span> | The name of the GKE project containing your prod cluster | <span style="font-size:0.75em;">your-project</span> | Only if you want to deploy on **GKE** from main/master branch to your prod cluster |
| <span style="font-size:0.75em;">LKE_SERVICEACCOUNT_TOKEN</span> | This is a token an `API Token` (also called `Personal Access Token`) created in a bot account created for service purposes | <span style="font-size:0.75em;">AR4nd0m5tr1ng</span> | Only if you want to deploy on **LKE** |
| <span style="font-size:0.75em;">LKE_DEV_CLUSTERID</span> | The number in your DEV Linode Kubernetes Cluster URL `/kubernetes/clusters/*****/summary` | <span style="font-size:0.75em;">99999</span> | Only if you want to deploy on **LKE** from dev branch to your dev cluster |
| <span style="font-size:0.75em;">LKE_PROD_CLUSTERID</span> | The number in your PROD Linode Kubernetes Cluster URL `/kubernetes/clusters/*****/summary` | <span style="font-size:0.75em;">99999</span> | Only if you want to deploy on **LKE** from main/master branch to your prod cluster |
### Add Deployment code snippet
You can add any number of the following snippets to the end of your push-{branch}.yml file to enable the deployment of the selected platform (Secrets must have been populated as describe above). An important step while injecting these code blocks is to make sure that the added job is tabbed like the other jobs already in the workflow file.
#### GKE's CD script
```yaml
  deployment-gke:
    needs: [publish-docker, publish-helm]
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Prepare deploy env
        run: |
            chmod 700 .github/tools_install.sh
            ./.github/tools_install.sh
            printf '${{ env.GKE_JSON }}' > deployments-data.json
            gcloud auth activate-service-account ${{ env.GKE_USER }} --key-file=deployments-data.json
            gcloud container clusters get-credentials ${{ env.GKE_CLUSTER_NAME }} --zone ${{ env.GKE_CLUSTER_ZONE }} --project ${{ env.GKE_CLUSTER_PROJECT }}

      - name: Push with valuefile
        run: |
            IN=$(echo ${GITHUB_REPOSITORY})
            NAME=${IN#"Ubivius/"}
            VERSION=$(yq e '.version' ./chart/Chart.yaml)
            helm upgrade $NAME --install --version $VERSION ysma500/$NAME -f chart/values.yaml
```
#### LKE's CD script
```yaml
  deployment-lke:
    needs: [publish-docker, publish-helm]
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Prepare deploy env
        run: |
            chmod 700 .github/tools_install.sh
            ./.github/tools_install.sh
            curl -H "Authorization: Bearer ${{ env.LKE_SERVICEACCOUNT_TOKEN }}" https://api.linode.com/v4/lke/clusters/${{ env.LKE_CLUSTERID }}/kubeconfig | jq '.kubeconfig' | sed 's/["]//g' | base64 --decode > ~/.kube/config

      - name: Push with valuefile
        run: |
            IN=$(echo ${GITHUB_REPOSITORY})
            NAME=${IN#"Ubivius/"}
            VERSION=$(yq e '.version' ./chart/Chart.yaml)
            helm upgrade $NAME --install --version $VERSION ysma500/$NAME -f chart/values.yaml
```
