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
import {
    TetrateApiTccCoreV1Tenant,
    TetrateApiTccCoreV1TenantFromJSON,
    TetrateApiTccCoreV1TenantToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1ListTenantsResponse
 */
export interface TetrateApiTccCoreV1ListTenantsResponse {
    /**
     * 
     * @type {Array<TetrateApiTccCoreV1Tenant>}
     * @memberof TetrateApiTccCoreV1ListTenantsResponse
     */
    tenants?: Array<TetrateApiTccCoreV1Tenant>;
}

export function TetrateApiTccCoreV1ListTenantsResponseFromJSON(json: any): TetrateApiTccCoreV1ListTenantsResponse {
    return {
        'tenants': !exists(json, 'tenants') ? undefined : (json['tenants'] as Array<any>).map(TetrateApiTccCoreV1TenantFromJSON),
    };
}

export function TetrateApiTccCoreV1ListTenantsResponseToJSON(value?: TetrateApiTccCoreV1ListTenantsResponse): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'tenants': value.tenants === undefined ? undefined : (value.tenants as Array<any>).map(TetrateApiTccCoreV1TenantToJSON),
    };
}


