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
 * @interface TetrateApiTccCoreV1RegistryConsumerResponse
 */
export interface TetrateApiTccCoreV1RegistryConsumerResponse {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1RegistryConsumerResponse
     */
    status?: string;
}

export function TetrateApiTccCoreV1RegistryConsumerResponseFromJSON(json: any): TetrateApiTccCoreV1RegistryConsumerResponse {
    return {
        'status': !exists(json, 'status') ? undefined : json['status'],
    };
}

export function TetrateApiTccCoreV1RegistryConsumerResponseToJSON(value?: TetrateApiTccCoreV1RegistryConsumerResponse): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'status': value.status,
    };
}


