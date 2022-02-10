> NOTE some step in this document might refers to Red Hat internal components that you do not have access to

# Populating Configuration

# Populating Required Configuration

1. Take note of your organization's `external_id` (your Red Hat Account's Organization ID). One way to retrieve
your organization's `external_id` is by logging in to OpenShift Cluster Manager (OCM) using its CLI tool: Log in (see https://github.com openshift-online/ocm-cli#log-in) and then run the `ocm whoami` command and take note of the `.organization.external_id` attribute
2. Add your organization's `external_id` (your Red Hat Account's Organization ID) to the [Quota Management List Configurations](quota-management-list-configuration.md) if you need to create STANDARD dinosaur instances. Follow the guide in [Quota Management List Configurations](access-control.md).
3. Follow the guide in [Access Control Configurations](access-control.md) to configure access control as required.

## Interacting with the Fleet Manager API

The Fleet Manager API requires OIDC Bearer tokens, which use the JWT format,
for authentication. The token is passed as an `Authorization` HTTP header value.

To be able to perform requests to the Fleet Manager API, you first need
to get a JWT token from the configured OIDC authentication server, which is
Red Hat SSO (sso.redhat.com) by default. Assuming the default OIDC authentication server is
being used, this can be performed by interacting with the OCM API. This can be
easily done interacting with it through the [ocm cli](https://github.com/openshift-online/ocm-cli/releases)
and retrieve OCM tokens using it. To do so:
1. Login to your desired OCM environment via web console using your Red Hat
   account credentials. For example, for the OCM production environment, go to
   https://cloud.redhat.com and login.
1. Get your OCM offline token by going to https://cloud.redhat.com/openshift/token
1. Login to your desired OCM environment through the OCM CLI by providing the
   OCM offline token and environment information:
   ```
   ocm login --token <ocm-offline-token> --url <ocm-api-url>
   ```
   `<ocm-api-url>` is the URL of the OCM API. Some shorthands can also
   be provided like `production` or `staging`
1. Generate an OCM token by running: `ocm token`. The generated token is the token
   that should be used to perform a request to the Fleet Manager API. For example:
  ```
  curl -H "Authorization: Bearer <result-of-ocm-token-command>" http://127.0.0.1:/8000/api/dinosaurs_mgmt
  ```
  OCM tokens other than the OCM offline token have an expiration time so a
  new one will need to be generated when that happens

There are some additional steps needed if you want to be able to perform
certain actions that have additional requirements. See the
[_User Account & Organization Setup_](getting-credentials-and-accounts.md#user-account--organization-setup) for more information

## Setting up OCM tokens

The Fleet Manager itself requires the use of an OCM token so it can
interact with OCM to perform management of Data Plane clusters.

In order for the Fleet Manager to do so, an OCM offline token should be configured.
To do so, retrieve your OCM offline token and then configure it for Fleet
Manager by running:
```
make ocm/setup OCM_OFFLINE_TOKEN=<your-retrieved-ocm-offline-token>
```

# Populating Optional Configuration
The subsections below describe how to setup some advanced configuration options.
To have a fully blown working service, a user is expected to setup each configuration.
However if you do not all the credentials you should still be able to run the fleet manager and interact with the public endpoints
to test things out.

>NOTE: See the [Makefile](../Makefile) for the environment variables accepted by each of the below commands.

>NOTE: You could pass dummy values to make sure that the secrets files are created.

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

## Setup the data plane image pull secret
In the Data Plane cluster, the Dinosaur Operator and the FleetShard Deployments
might reference container images that are located in authenticated container
image registries.

Fleet Manager can be configured to send this authenticated
container image registry information as a K8s Secret in [`kubernetes.io/.dockerconfigjson` format](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/#registry-secret-existing-credentials).

In order for the Fleet Manager to be able to start, create the following file:
```
touch secrets/image-pull.dockerconfigjson
```

If you don't need to make use of this functionality you can skip this section.
Otherwise, keep reading below.

To configure the Fleet Manager with this authenticated registry information so
the previously mentioned Data Plane elements can pull container images from it:
* Base-64 encode your [Docker configuration file](https://docs.docker.com/engine/reference/commandline/cli/#docker-cli-configuration-file-configjson-properties).
* Copy the contents generated from the previous point into the `secrets/image-pull.dockerconfigjson` file

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
