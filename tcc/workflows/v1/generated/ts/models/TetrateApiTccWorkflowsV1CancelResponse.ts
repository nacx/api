// tslint:disable
/**
 * loadbalancer.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists } from '../runtime';
/**
 * 
 * @export
 * @interface TetrateApiTccWorkflowsV1CancelResponse
 */
export interface TetrateApiTccWorkflowsV1CancelResponse {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1CancelResponse
     */
    requestid?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1CancelResponse
     */
    note?: string;
}

export function TetrateApiTccWorkflowsV1CancelResponseFromJSON(json: any): TetrateApiTccWorkflowsV1CancelResponse {
    return {
        'requestid': !exists(json, 'requestid') ? undefined : json['requestid'],
        'note': !exists(json, 'note') ? undefined : json['note'],
    };
}

export function TetrateApiTccWorkflowsV1CancelResponseToJSON(value?: TetrateApiTccWorkflowsV1CancelResponse): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'requestid': value.requestid,
        'note': value.note,
    };
}


