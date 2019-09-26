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
    TetrateApiTccCoreV1Team,
    TetrateApiTccCoreV1TeamFromJSON,
    TetrateApiTccCoreV1TeamToJSON,
} from './';

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
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateTeamRequest
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateTeamRequest
     */
    tenant?: string;
    /**
     * 
     * @type {TetrateApiTccCoreV1Team}
     * @memberof TetrateApiTccCoreV1CreateTeamRequest
     */
    team?: TetrateApiTccCoreV1Team;
}

export function TetrateApiTccCoreV1CreateTeamRequestFromJSON(json: any): TetrateApiTccCoreV1CreateTeamRequest {
    return {
        'parent': !exists(json, 'parent') ? undefined : json['parent'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'team': !exists(json, 'team') ? undefined : TetrateApiTccCoreV1TeamFromJSON(json['team']),
    };
}

export function TetrateApiTccCoreV1CreateTeamRequestToJSON(value?: TetrateApiTccCoreV1CreateTeamRequest): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'parent': value.parent,
        'id': value.id,
        'tenant': value.tenant,
        'team': TetrateApiTccCoreV1TeamToJSON(value.team),
    };
}


