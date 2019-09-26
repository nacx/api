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
import {
    TetrateApiTccCoreV1Permissions,
    TetrateApiTccCoreV1PermissionsFromJSON,
    TetrateApiTccCoreV1PermissionsToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1UpdateApplicationRequest
 */
export interface TetrateApiTccCoreV1UpdateApplicationRequest {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1UpdateApplicationRequest
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1UpdateApplicationRequest
     */
    displayName?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1UpdateApplicationRequest
     */
    workspace?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1UpdateApplicationRequest
     */
    tenant?: string;
    /**
     * 
     * @type {TetrateApiTccCoreV1Permissions}
     * @memberof TetrateApiTccCoreV1UpdateApplicationRequest
     */
    permissions?: TetrateApiTccCoreV1Permissions;
}

export function TetrateApiTccCoreV1UpdateApplicationRequestFromJSON(json: any): TetrateApiTccCoreV1UpdateApplicationRequest {
    return {
        'name': !exists(json, 'name') ? undefined : json['name'],
        'displayName': !exists(json, 'displayName') ? undefined : json['displayName'],
        'workspace': !exists(json, 'workspace') ? undefined : json['workspace'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'permissions': !exists(json, 'permissions') ? undefined : TetrateApiTccCoreV1PermissionsFromJSON(json['permissions']),
    };
}

export function TetrateApiTccCoreV1UpdateApplicationRequestToJSON(value?: TetrateApiTccCoreV1UpdateApplicationRequest): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'name': value.name,
        'displayName': value.displayName,
        'workspace': value.workspace,
        'tenant': value.tenant,
        'permissions': TetrateApiTccCoreV1PermissionsToJSON(value.permissions),
    };
}


