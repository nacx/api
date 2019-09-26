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
 * @interface TetrateApiTccCoreV1CreateTenantRequest
 */
export interface TetrateApiTccCoreV1CreateTenantRequest {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateTenantRequest
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateTenantRequest
     */
    description?: string;
}

export function TetrateApiTccCoreV1CreateTenantRequestFromJSON(json: any): TetrateApiTccCoreV1CreateTenantRequest {
    return {
        'id': !exists(json, 'id') ? undefined : json['id'],
        'description': !exists(json, 'description') ? undefined : json['description'],
    };
}

export function TetrateApiTccCoreV1CreateTenantRequestToJSON(value?: TetrateApiTccCoreV1CreateTenantRequest): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'id': value.id,
        'description': value.description,
    };
}


