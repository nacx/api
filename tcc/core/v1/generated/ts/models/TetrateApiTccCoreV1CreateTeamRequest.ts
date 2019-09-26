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
 * @interface TetrateApiTccCoreV1CreateTeamRequest
 */
export interface TetrateApiTccCoreV1CreateTeamRequest {
    /**
     * Internal use only. Auto populated field.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateTeamRequest
     */
    parent?: string;
    /**
     * Tenant.Id
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateTeamRequest
     */
    tenant?: string;
    /**
     * if present, this will be used as the id for the created object.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateTeamRequest
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateTeamRequest
     */
    description?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof TetrateApiTccCoreV1CreateTeamRequest
     */
    members?: Array<string>;
}

export function TetrateApiTccCoreV1CreateTeamRequestFromJSON(json: any): TetrateApiTccCoreV1CreateTeamRequest {
    return {
        'parent': !exists(json, 'parent') ? undefined : json['parent'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'members': !exists(json, 'members') ? undefined : json['members'],
    };
}

export function TetrateApiTccCoreV1CreateTeamRequestToJSON(value?: TetrateApiTccCoreV1CreateTeamRequest): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'parent': value.parent,
        'tenant': value.tenant,
        'id': value.id,
        'description': value.description,
        'members': value.members,
    };
}


