> NOTE some step in this document might refers to Red Hat internal components that you do not have access to
 
# Populating Configuration
1. Add your organization's `external_id` to the [Quota Management List Configurations](quota-management-list-configuration.md) 
if you need to create STANDARD dinosaur instances. Follow the guide in [Quota Management List Configurations](access-control.md)
2. Follow the guide in [Access Control Configurations](access-control.md) to configure access control as required.
3. Retrieve your ocm-offline-token from https://qaprodauth.cloud.redhat.com/openshift/token and save it to `secrets/ocm-service.token` 
4. Setup AWS configuration
```
make aws/setup
```
5. Setup SSO configuration
    - mas sso client id & client secret
    ```
    make keycloak/setup SSO_CLIENT_ID=<SSO_client_id> SSO_CLIENT_SECRET=<SSO_client_secret> OSD_IDP_SSO_CLIENT_ID=<osd_idp_SSO_client_id> OSD_IDP_SSO_CLIENT_SECRET=<osd_idp_SSO_client_secret>
    ```
6. Setup Dinosaur TLS cert
```
make dinosaurcert/setup
```
7. Setup the image pull secret
    - To be able to pull docker images from quay, copy the content of your account's secret (`config.json` key) and paste it to `secrets/image-pull.dockerconfigjson` file.

8. Setup the Observability stack secrets
```
make observatorium/setup
```
