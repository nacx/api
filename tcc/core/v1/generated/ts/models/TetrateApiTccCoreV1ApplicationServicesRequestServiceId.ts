// tslint:disable
/**
 * TCC Core API
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
 * @interface TetrateApiTccCoreV1ApplicationServicesRequestServiceId
 */
export interface TetrateApiTccCoreV1ApplicationServicesRequestServiceId {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1ApplicationServicesRequestServiceId
     */
    cluster?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1ApplicationServicesRequestServiceId
     */
    namespace?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1ApplicationServicesRequestServiceId
     */
    hostname?: string;
}

export function TetrateApiTccCoreV1ApplicationServicesRequestServiceIdFromJSON(json: any): TetrateApiTccCoreV1ApplicationServicesRequestServiceId {
    return {
        'cluster': !exists(json, 'cluster') ? undefined : json['cluster'],
        'namespace': !exists(json, 'namespace') ? undefined : json['namespace'],
        'hostname': !exists(json, 'hostname') ? undefined : json['hostname'],
    };
}

export function TetrateApiTccCoreV1ApplicationServicesRequestServiceIdToJSON(value?: TetrateApiTccCoreV1ApplicationServicesRequestServiceId): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'cluster': value.cluster,
        'namespace': value.namespace,
        'hostname': value.hostname,
    };
}


