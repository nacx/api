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
 * @interface TetrateApiTccCoreV1CreateUserRequest
 */
export interface TetrateApiTccCoreV1CreateUserRequest {
    /**
     * Internal use only. Auto populated field.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateUserRequest
     */
    parent?: string;
    /**
     * Tenant.Id
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateUserRequest
     */
    tenant?: string;
    /**
     * if present, this will be used as the id for the created object.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateUserRequest
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateUserRequest
     */
    description?: string;
}

export function TetrateApiTccCoreV1CreateUserRequestFromJSON(json: any): TetrateApiTccCoreV1CreateUserRequest {
    return {
        'parent': !exists(json, 'parent') ? undefined : json['parent'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'description': !exists(json, 'description') ? undefined : json['description'],
    };
}

export function TetrateApiTccCoreV1CreateUserRequestToJSON(value?: TetrateApiTccCoreV1CreateUserRequest): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'parent': value.parent,
        'tenant': value.tenant,
        'id': value.id,
        'description': value.description,
    };
}


