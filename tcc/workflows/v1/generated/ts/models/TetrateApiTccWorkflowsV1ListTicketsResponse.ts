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
    TetrateApiTccWorkflowsV1ListTicketsResponsePendingTickets,
    TetrateApiTccWorkflowsV1ListTicketsResponsePendingTicketsFromJSON,
    TetrateApiTccWorkflowsV1ListTicketsResponsePendingTicketsToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccWorkflowsV1ListTicketsResponse
 */
export interface TetrateApiTccWorkflowsV1ListTicketsResponse {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1ListTicketsResponse
     */
    cluster?: string;
    /**
     * 
     * @type {Array<TetrateApiTccWorkflowsV1ListTicketsResponsePendingTickets>}
     * @memberof TetrateApiTccWorkflowsV1ListTicketsResponse
     */
    pendingTickets?: Array<TetrateApiTccWorkflowsV1ListTicketsResponsePendingTickets>;
}

export function TetrateApiTccWorkflowsV1ListTicketsResponseFromJSON(json: any): TetrateApiTccWorkflowsV1ListTicketsResponse {
    return {
        'cluster': !exists(json, 'cluster') ? undefined : json['cluster'],
        'pendingTickets': !exists(json, 'pendingTickets') ? undefined : (json['pendingTickets'] as Array<any>).map(TetrateApiTccWorkflowsV1ListTicketsResponsePendingTicketsFromJSON),
    };
}

export function TetrateApiTccWorkflowsV1ListTicketsResponseToJSON(value?: TetrateApiTccWorkflowsV1ListTicketsResponse): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'cluster': value.cluster,
        'pendingTickets': value.pendingTickets === undefined ? undefined : (value.pendingTickets as Array<any>).map(TetrateApiTccWorkflowsV1ListTicketsResponsePendingTicketsToJSON),
    };
}


