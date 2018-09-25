# machine-controller-manager-provider-aws

- This repository aims to implement the out-of-tree driver support for AWS in [machine-controller-manager](https://github.com/gardener/machine-controller-manager/tree/grpc-driver)
- This makes use of gRPC to achieve the communication between the Driver and MCM.

## Usage


#### Deploy the required CRDs
- Deploy the required CRDs for the machine-controller-manager by pointing your KUBECONFIG to the appropriate cluster.
```
    kubectl apply -f kubernetes/crds.yaml
```

#### Starting the machine-controller-manager
- Fetch the [machine-controller-manager](https://github.com/gardener/machine-controller-manager/tree/grpc-driver) repository and cd into it.
```
     git clone https://github.com/gardener/machine-controller-manager/
```
- Now checkout into the grpc-driver branch.
```
    git checkout grpc-driver
```
- Make sure your `Makefile` is pointing to the right clusters.
- Start the machine-controller-manager.
```
    make start-with-grpc
```

#### Starting the machine-controller-manager-provider-aws
- Now open another terminal and fetch [this repository](/) and cd into it.
- Start the aws driver.
```
    make start-with-aws
```

#### Creating your own machine
- Edit the machine-class file - s`kubernetes/aws-machine-class.yaml` and replace the necessary details. And apply the same by pointing your KUBECONFIG to the right cluster.
```
    kubectl apply -f kubernetes/aws-machine-class.yaml
```
- Edit the secret `kubernetes/aws-secret.yaml` file and replace the necessary details.
```
    kubectl apply -f kubernetes/aws-secret.yaml
```
- Finally deploy the machine  `kubernetes/machine.yaml` by replacing any necessary details.
```
    kubectl apply -f kubernetes/aws-secret.yaml
```

## Sequence of steps that would occur during machine creation

<img width="1254" alt="screen shot 2018-09-20 at 2 12 04 pm" src="https://user-images.githubusercontent.com/31065672/45806640-35bb5b00-bcdf-11e8-9811-f8e258416bf6.png">
