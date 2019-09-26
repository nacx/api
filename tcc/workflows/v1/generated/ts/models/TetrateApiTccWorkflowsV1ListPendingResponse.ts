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
import {
    TetrateApiTccWorkflowsV1ListPendingResponsePendingRequest,
    TetrateApiTccWorkflowsV1ListPendingResponsePendingRequestFromJSON,
    TetrateApiTccWorkflowsV1ListPendingResponsePendingRequestToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccWorkflowsV1ListPendingResponse
 */
export interface TetrateApiTccWorkflowsV1ListPendingResponse {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1ListPendingResponse
     */
    cluster?: string;
    /**
     * 
     * @type {Array<TetrateApiTccWorkflowsV1ListPendingResponsePendingRequest>}
     * @memberof TetrateApiTccWorkflowsV1ListPendingResponse
     */
    pendingRequests?: Array<TetrateApiTccWorkflowsV1ListPendingResponsePendingRequest>;
}

export function TetrateApiTccWorkflowsV1ListPendingResponseFromJSON(json: any): TetrateApiTccWorkflowsV1ListPendingResponse {
    return {
        'cluster': !exists(json, 'cluster') ? undefined : json['cluster'],
        'pendingRequests': !exists(json, 'pendingRequests') ? undefined : (json['pendingRequests'] as Array<any>).map(TetrateApiTccWorkflowsV1ListPendingResponsePendingRequestFromJSON),
    };
}

export function TetrateApiTccWorkflowsV1ListPendingResponseToJSON(value?: TetrateApiTccWorkflowsV1ListPendingResponse): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'cluster': value.cluster,
        'pendingRequests': value.pendingRequests === undefined ? undefined : (value.pendingRequests as Array<any>).map(TetrateApiTccWorkflowsV1ListPendingResponsePendingRequestToJSON),
    };
}


