> NOTE some step in this document might refers to Red Hat internal components that you do not have access to
 
# Populating Configuration

# Populating Required Configuration
1. Take note of your organization's `external_id` (your Red Hat Account's Organization ID). One way to retrieve
your organization's `external_id` is by logging in to OpenShift Cluster Manager (OCM) using its CLI tool: Log in (see https://github.com openshift-online/ocm-cli#log-in) and then run the `ocm whoami` command and take note of the `.organization.external_id` attribute
2. Add your organization's `external_id` (your Red Hat Account's Organization ID) to the [Quota Management List Configurations](quota-management-list-configuration.md) if you need to create STANDARD dinosaur instances. Follow the guide in [Quota Management List Configurations](access-control.md).
3. Follow the guide in [Access Control Configurations](access-control.md) to configure access control as required.

# Populating Optional Configuration
The subsections below describe how to setup some advanced configuration options.
To have a fully blown working service, a user is expected to setup each configuration. 
However if you do not all the credentials you should still be able to run the fleet manager and interact with the public endpoints
to test things out.

>NOTE: See the [Makefile](../Makefile) for the environment variables accepted by each of the below commands.

>NOTE: You could pass dummy values to make sure that the secrets files are created.

## Setting up OCM token

[OpenShift Cluster Managament](https://api.openshift.com/) is utilised for the management of Data plane clusters. 
To be able to talk to this service, retrieve your ocm-offline-token from https://qaprodauth.cloud.redhat.com/openshift/token 
and save and run the following command
```
make ocm/setup OCM_OFFLINE_TOKEN=<your-retrieved-offline-token>
```

See the [_User Account & Organization Setup_](getting-credentials-and-accounts.md#user-account--organization-setup) for info on how to get access to the api.

## Setup AWS configuration
This is used to initialise the aws accounts for cluster creation via OCM and also to setup AWS route53 account.
```
make aws/setup
```

## Setup SSO configuration
mas sso is used to create a Keycloak client used for communication between the fleetshard operator and the fleet manager.
Additionally, it is used to configure the OpenShift cluster identity provider to give SRE an access to the cluster.
```
make keycloak/setup
```
>NOTE: MAS-SSO is soon going to be deprecated. Reach out to the MAS-Security team for info regarding this.
Alternatively, you could run your own Keycloak instance by following their [getting-started guide](https://www.keycloak.org/getting-started/getting-started-docker). Please note, how to setup Keycloak is out of scope of this guide

## Setup the image pull secret
To be able to pull docker images from quay, copy the content of your account's secret (`config.json` key) and paste it to `secrets/image-pull.dockerconfigjson` file.

## Setup the Observability stack secrets
See [Obsevability](./observability/README.md) to learn more about Observatorium and the observability stack.
The following command is used to setup the various secrets needed by the Observability stack.

```
make observatorium/setup
```

## Setup Dinosaur Service TLS
Setup the dinosaur TLS certificate used for Managed Dinosaur Service
```
make dinosaurcert/setup
```
