## AllDayDevOps Spring Break Edition 2020

Cloud Native Jenkins cluster deployment on Azure with Terraform & Packer. Plus, setup of CI/CD pipeline on Jenkins with Pipeline as Code.

## Jenkins Workers Resource Identity

- Create Managed Identity called JenkinsWorker:

```
az identity create --resource-group $resourceGroup --name JenkinsWorker
```

- Assign JenkinsWorker identity to VMs

- Get JenkinsWorker Principal ID:

```
principalId=$(az identity show --resource-group $resourceGroup --name JenkinsWorker --query principalId --output tsv)
```

- Assign the desired role the managed identity:

```
containerId=$(az container show --resource-group $resourceGroup --name application-sandbox|staging|production --query id --output tsv)
az role assignment create --assignee $principalId --scope $containerId --role Contributor
```

## Author 

Mohamed Labouardy