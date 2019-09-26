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
 * @interface TetrateApiTccWorkflowsV1ApproveResponse
 */
export interface TetrateApiTccWorkflowsV1ApproveResponse {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1ApproveResponse
     */
    requestid?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1ApproveResponse
     */
    note?: string;
}

export function TetrateApiTccWorkflowsV1ApproveResponseFromJSON(json: any): TetrateApiTccWorkflowsV1ApproveResponse {
    return {
        'requestid': !exists(json, 'requestid') ? undefined : json['requestid'],
        'note': !exists(json, 'note') ? undefined : json['note'],
    };
}

export function TetrateApiTccWorkflowsV1ApproveResponseToJSON(value?: TetrateApiTccWorkflowsV1ApproveResponse): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'requestid': value.requestid,
        'note': value.note,
    };
}


