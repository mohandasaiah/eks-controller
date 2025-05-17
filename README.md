# EKS Controller 🚀

![GitHub release](https://img.shields.io/github/release/mohandasaiah/eks-controller.svg)
![GitHub issues](https://img.shields.io/github/issues/mohandasaiah/eks-controller.svg)

Welcome to the EKS Controller repository! This project provides an ACK (AWS Controllers for Kubernetes) service controller for Amazon Elastic Kubernetes Service (EKS). The EKS Controller simplifies the management of AWS resources directly from your Kubernetes cluster.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Introduction

The EKS Controller enables you to manage AWS services using Kubernetes-native resources. This integration allows developers to use Kubernetes APIs to interact with AWS services, enhancing automation and efficiency. 

For the latest releases, visit [this link](https://github.com/mohandasaiah/eks-controller/releases). You can download the necessary files and execute them as needed.

## Features

- **Seamless Integration**: Easily manage AWS resources from within your Kubernetes cluster.
- **Declarative Configuration**: Use Kubernetes manifests to define and manage your AWS resources.
- **Extensible**: Add new AWS services and features as needed.
- **Community Driven**: Contributions from the community enhance functionality and stability.

## Installation

To get started with the EKS Controller, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/mohandasaiah/eks-controller.git
   cd eks-controller
   ```

2. **Install Dependencies**:
   Ensure you have the necessary tools installed, such as `kubectl` and `AWS CLI`.

3. **Deploy the Controller**:
   Apply the deployment manifests to your EKS cluster:
   ```bash
   kubectl apply -f deploy/
   ```

4. **Verify Installation**:
   Check if the EKS Controller is running:
   ```bash
   kubectl get pods -n eks-controller
   ```

For the latest releases, visit [this link](https://github.com/mohandasaiah/eks-controller/releases). You can download the necessary files and execute them as needed.

## Usage

After installation, you can start using the EKS Controller to manage AWS resources. Here’s a simple example of how to create an S3 bucket:

1. **Create a Manifest**:
   Create a file named `s3-bucket.yaml` with the following content:
   ```yaml
   apiVersion: s3.services.k8s.aws/v1alpha1
   kind: Bucket
   metadata:
     name: my-s3-bucket
   spec:
     location: us-west-2
   ```

2. **Apply the Manifest**:
   Run the following command to create the S3 bucket:
   ```bash
   kubectl apply -f s3-bucket.yaml
   ```

3. **Check the Status**:
   Verify the bucket creation:
   ```bash
   kubectl get buckets
   ```

## Configuration

You can configure the EKS Controller to suit your needs. The configuration can be done through Kubernetes manifests. Here are some common configurations:

- **IAM Roles**: Ensure that your Kubernetes service account has the necessary IAM roles attached to access AWS resources.
- **Region Configuration**: Specify the AWS region in your manifests to ensure resources are created in the correct location.

## Contributing

We welcome contributions to the EKS Controller! If you would like to contribute, please follow these steps:

1. **Fork the Repository**: Create your own fork of the repository.
2. **Create a Branch**: Create a new branch for your feature or bug fix.
   ```bash
   git checkout -b feature/my-feature
   ```
3. **Make Changes**: Implement your changes and ensure they are well-tested.
4. **Submit a Pull Request**: Push your changes and submit a pull request for review.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, please reach out to the project maintainer:

- **Name**: Mohan Dasaiah
- **Email**: mohandasaiah@example.com

Thank you for your interest in the EKS Controller! We hope you find it useful for managing your AWS resources through Kubernetes.