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

/**
 * 
 * @export
 * @enum {string}
 */
export enum TetrateApiTccWorkflowsV1ListPendingResponsePendingRequestOP {
    ATTACH = 'ATTACH',
    DETACH = 'DETACH'
}

export function TetrateApiTccWorkflowsV1ListPendingResponsePendingRequestOPFromJSON(json: any): TetrateApiTccWorkflowsV1ListPendingResponsePendingRequestOP {
    return json as TetrateApiTccWorkflowsV1ListPendingResponsePendingRequestOP;
}

export function TetrateApiTccWorkflowsV1ListPendingResponsePendingRequestOPToJSON(value?: TetrateApiTccWorkflowsV1ListPendingResponsePendingRequestOP): any {
    return value as any;
}

