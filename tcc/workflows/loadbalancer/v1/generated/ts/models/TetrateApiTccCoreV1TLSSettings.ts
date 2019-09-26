// tslint:disable
/**
 * TCC Workflows API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: v1
 * Contact: info@tetrate.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1TLSSettings
 */
export interface TetrateApiTccCoreV1TLSSettings {
    /**
     * enable TLS settings for the application.
     * @type {boolean}
     * @memberof TetrateApiTccCoreV1TLSSettings
     */
    enabled?: boolean;
    /**
     * If set, the load balancer will redirect HTTP connections on port 80 to HTTPS port 443.
     * @type {boolean}
     * @memberof TetrateApiTccCoreV1TLSSettings
     */
    redirectToHttps?: boolean;
    /**
     * For proxies running on VMs, the path to the file holding the server-side TLS certificate to use.
     * @type {string}
     * @memberof TetrateApiTccCoreV1TLSSettings
     */
    serverCertificate?: string;
    /**
     * For proxies running on VMs, the path to the file holding the server\'s private key.
     * @type {string}
     * @memberof TetrateApiTccCoreV1TLSSettings
     */
    privateKey?: string;
    /**
     * For proxies running on VMs, the path to a file containing certificate authority certificates to use in verifying a presented client side certificate for mutual TLS connections.
     * @type {string}
     * @memberof TetrateApiTccCoreV1TLSSettings
     */
    caCertificates?: string;
    /**
     * For proxies running on Kubernetes, the name of the secret that holds the TLS certs. Currently applicable only on Kubernetes. The secret should contain the server certificate and the private key. If mutual TLS is being used, an additional secret with name secretName-cacert should be created with the CaCertificates that the server will use to verify client side certificates. If the service is exposed via a load balancer, the secret must be accessible to it.
     * @type {string}
     * @memberof TetrateApiTccCoreV1TLSSettings
     */
    secretName?: string;
}

export function TetrateApiTccCoreV1TLSSettingsFromJSON(json: any): TetrateApiTccCoreV1TLSSettings {
    return {
        'enabled': !exists(json, 'enabled') ? undefined : json['enabled'],
        'redirectToHttps': !exists(json, 'redirectToHttps') ? undefined : json['redirectToHttps'],
        'serverCertificate': !exists(json, 'serverCertificate') ? undefined : json['serverCertificate'],
        'privateKey': !exists(json, 'privateKey') ? undefined : json['privateKey'],
        'caCertificates': !exists(json, 'caCertificates') ? undefined : json['caCertificates'],
        'secretName': !exists(json, 'secretName') ? undefined : json['secretName'],
    };
}

export function TetrateApiTccCoreV1TLSSettingsToJSON(value?: TetrateApiTccCoreV1TLSSettings): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'enabled': value.enabled,
        'redirectToHttps': value.redirectToHttps,
        'serverCertificate': value.serverCertificate,
        'privateKey': value.privateKey,
        'caCertificates': value.caCertificates,
        'secretName': value.secretName,
    };
}


