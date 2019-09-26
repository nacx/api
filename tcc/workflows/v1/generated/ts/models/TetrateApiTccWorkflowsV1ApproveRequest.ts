// tslint:disable
/**
 * TCC Workflows API
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
 * @interface TetrateApiTccWorkflowsV1ApproveRequest
 */
export interface TetrateApiTccWorkflowsV1ApproveRequest {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1ApproveRequest
     */
    requestid?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1ApproveRequest
     */
    cluster?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1ApproveRequest
     */
    tenant?: string;
}

export function TetrateApiTccWorkflowsV1ApproveRequestFromJSON(json: any): TetrateApiTccWorkflowsV1ApproveRequest {
    return {
        'requestid': !exists(json, 'requestid') ? undefined : json['requestid'],
        'cluster': !exists(json, 'cluster') ? undefined : json['cluster'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
    };
}

export function TetrateApiTccWorkflowsV1ApproveRequestToJSON(value?: TetrateApiTccWorkflowsV1ApproveRequest): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'requestid': value.requestid,
        'cluster': value.cluster,
        'tenant': value.tenant,
    };
}


