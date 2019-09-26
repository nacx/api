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
 * @interface TetrateApiTccWorkflowsV1LBTicketResolution
 */
export interface TetrateApiTccWorkflowsV1LBTicketResolution {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1LBTicketResolution
     */
    tenant?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1LBTicketResolution
     */
    workspace?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1LBTicketResolution
     */
    requestid?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccWorkflowsV1LBTicketResolution
     */
    note?: string;
}

export function TetrateApiTccWorkflowsV1LBTicketResolutionFromJSON(json: any): TetrateApiTccWorkflowsV1LBTicketResolution {
    return {
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'workspace': !exists(json, 'workspace') ? undefined : json['workspace'],
        'requestid': !exists(json, 'requestid') ? undefined : json['requestid'],
        'note': !exists(json, 'note') ? undefined : json['note'],
    };
}

export function TetrateApiTccWorkflowsV1LBTicketResolutionToJSON(value?: TetrateApiTccWorkflowsV1LBTicketResolution): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'tenant': value.tenant,
        'workspace': value.workspace,
        'requestid': value.requestid,
        'note': value.note,
    };
}


