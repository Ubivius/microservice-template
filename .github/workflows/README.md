# Supported Cloud for Deployment
In Ubivius deployment solution, we are currently supporting these cloud providers:
- Google Kubernetes Engine(GKE)
- Linode Kubernetes Engine(LKE)

To customize you can add the desired deployment code block at the end of your `push-{target_branch}.yml` GitHub Action file following the procedure described below.

## Config deployment steps
### Add GitHub Secrets
<table>
  <thead>
    <tr>
      <th><sub>Secret Name</sub></th>
      <th><sub>Description</sub></th>
      <th><sub>Example</sub></th>
      <th><sub>Is it mandatory?</sub></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><sub>DOCKERHUB_USERNAME</sub></td>
      <td><sub>Your DockerHub service account short username used to push your built Docker container to DockerHub</sub></td>
      <td><sub>your_name</sub></td>
      <td><sub>Yes</sub></td>
    </tr>
    <tr>
      <td><sub>DOCKERHUB_PWD</sub></td>
      <td><sub>Your DockerHub service account's password to enable pushing built containers</sub></td>
      <td><sub>abc123</sub></td>
      <td><sub>Yes</sub></td>
    </tr>
    <tr>
      <td><sub>GKE_SERVICEACCOUNT_EMAIL</sub></td>
      <td><sub>This is your GKE service account's email</sub></td>
      <td><sub>my_deployment_bot@cluster_name.iam.gserviceaccount.com</sub></td>
      <td><sub>Only if you want to deploy on <b>GKE</b></sub></td>
    </tr>
    <tr>
      <td><sub>GKE_SERVICEACCOUNT_JSONKEY</sub></td>
      <td><sub>The downloaded json file's content when you create your service account's key</sub></td>
      <td><sub>{<br/>  "type": "{service_account}",<br/>  "project_id": "{project_id}",<br/>    "private_key_id": "{key_id}",<br/>  "private_key": "-----BEGIN PRIVATE KEY-----\{your_key}\n-----END PRIVATE KEY-----\n",<br/>  "client_email": "{service_account}",<br/>  "client_id": "{user_id}",<br/>  "auth_uri": "https://accounts.google.com/o/oauth2/auth",<br/>  "token_uri": "https://oauth2.googleapis.com/token",<br/>  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",<br/>  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/{service_account}"<br/>}</sub></td>
      <td><sub>Only if you want to deploy on <b>GKE</b></sub></td>
    </tr>
    <tr>
      <td><sub>GKE_DEV_CLUSTER_NAME</sub></td>
      <td><sub>The name of your dev cluster</sub></td>
      <td><sub>dev-environment</sub></td>
      <td><sub>Only if you want to deploy on <b>GKE</b> from dev branch to your dev cluster</sub></td>
    </tr>
    <tr>
      <td><sub>GKE_DEV_CLUSTER_ZONE</sub></td>
      <td><sub>The zone your dev cluster is located</sub></td>
      <td><sub>us-east1-c</sub></td>
      <td><sub>Only if you want to deploy on <b>GKE</b> from dev branch to your dev cluster</sub></td>
    </tr>
    <tr>
      <td><sub>GKE_DEV_CLUSTER_PROJECT</sub></td>
      <td><sub>The name of the GKE project containing your dev cluster</sub></td>
      <td><sub>your-project</sub></td>
      <td><sub>Only if you want to deploy on <b>GKE</b> from dev branch to your dev cluster</sub></td>
    </tr>
    <tr>
      <td><sub>GKE_PROD_CLUSTER_NAME</sub></td>
      <td><sub>The name of your prod cluster</sub></td>
      <td><sub>prod-environment</sub></td>
      <td><sub>Only if you want to deploy on <b>GKE</b> from main/master branch to your prod cluster</sub></td>
    </tr>
    <tr>
      <td><sub>GKE_PROD_CLUSTER_ZONE</sub></td>
      <td><sub>The zone your prod cluster is located</sub></td>
      <td><sub>us-east1-c</sub></td>
      <td><sub>Only if you want to deploy on <b>GKE</b> from main/master branch to your prod cluster</sub></td>
    </tr>
    <tr>
      <td><sub>GKE_PROD_CLUSTER_PROJECT</sub></td>
      <td><sub>The name of the GKE project containing your prod cluster</sub></td>
      <td><sub>your-project</sub></td>
      <td><sub>Only if you want to deploy on <b>GKE</b> from main/master branch to your prod cluster</sub></td>
    </tr>
    <tr>
      <td><sub>LKE_SERVICEACCOUNT_TOKEN</sub></td>
      <td><sub>This is a token an `API Token` (also called `Personal Access Token`) created in a bot account created for service purposes</sub></td>
      <td><sub>AR4nd0m5tr1ng</sub></td>
      <td><sub>Only if you want to deploy on <b>LKE</b></sub></td>
    </tr>
    <tr>
      <td><sub>LKE_DEV_CLUSTERID</sub></td>
      <td><sub>The number in your DEV Linode Kubernetes Cluster URL clusters/*****/summary</sub></td>
      <td><sub>99999</sub></td>
      <td><sub>Only if you want to deploy on <b>LKE</b> from dev branch to your dev cluster</sub></td>
    </tr>
    <tr>
      <td><sub>LKE_PROD_CLUSTERID</sub></td>
      <td><sub>The number in your PROD Linode Kubernetes Cluster URL clusters/*****/summary</sub></td>
      <td><sub>99999</sub></td>
      <td><sub>Only if you want to deploy on <b>LKE</b> from main/master branch to your prod cluster</sub></td>
    </tr>
  </tbody>
</table>

### Add Deployment code snippet
You can add any number of the following snippets to the end of your push-{branch}.yml file to enable the deployment of the selected platform (Secrets must have been populated as describe above). An important step while injecting these code blocks is to make sure that the added job is tabbed like the other jobs already in the workflow file.
#### GKE's deployment script
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
#### LKE's deployment script
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
