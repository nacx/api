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
    TetrateApiTccCoreV1Team,
    TetrateApiTccCoreV1TeamFromJSON,
    TetrateApiTccCoreV1TeamToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1ListTeamsResponse
 */
export interface TetrateApiTccCoreV1ListTeamsResponse {
    /**
     * 
     * @type {Array<TetrateApiTccCoreV1Team>}
     * @memberof TetrateApiTccCoreV1ListTeamsResponse
     */
    teams?: Array<TetrateApiTccCoreV1Team>;
}

export function TetrateApiTccCoreV1ListTeamsResponseFromJSON(json: any): TetrateApiTccCoreV1ListTeamsResponse {
    return {
        'teams': !exists(json, 'teams') ? undefined : (json['teams'] as Array<any>).map(TetrateApiTccCoreV1TeamFromJSON),
    };
}

export function TetrateApiTccCoreV1ListTeamsResponseToJSON(value?: TetrateApiTccCoreV1ListTeamsResponse): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'teams': value.teams === undefined ? undefined : (value.teams as Array<any>).map(TetrateApiTccCoreV1TeamToJSON),
    };
}

