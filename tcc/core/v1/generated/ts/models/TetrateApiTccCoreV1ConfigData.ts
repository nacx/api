// tslint:disable
/**
 * TCC Configuration API
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
 * @interface TetrateApiTccCoreV1ConfigData
 */
export interface TetrateApiTccCoreV1ConfigData {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1ConfigData
     */
    configtype?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1ConfigData
     */
    cluster?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1ConfigData
     */
    environment?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1ConfigData
     */
    tenant?: string;
    /**
     * For istio, this is a huge YAML dump of all Istio artifacts like virtual services, destination rules, service entries, gateways, etc. in YAML form. For F5, its all f5 config artifacts.
     * @type {string}
     * @memberof TetrateApiTccCoreV1ConfigData
     */
    payload?: string;
}

export function TetrateApiTccCoreV1ConfigDataFromJSON(json: any): TetrateApiTccCoreV1ConfigData {
    return {
        'configtype': !exists(json, 'configtype') ? undefined : json['configtype'],
        'cluster': !exists(json, 'cluster') ? undefined : json['cluster'],
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'payload': !exists(json, 'payload') ? undefined : json['payload'],
    };
}

export function TetrateApiTccCoreV1ConfigDataToJSON(value?: TetrateApiTccCoreV1ConfigData): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'configtype': value.configtype,
        'cluster': value.cluster,
        'environment': value.environment,
        'tenant': value.tenant,
        'payload': value.payload,
    };
}

