// tslint:disable
/**
 * configproducer.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: version not set
 * 
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
export enum TetrateApiTccV1WorkflowsLBClass {
    Envoy = 'Envoy',
    F5 = 'F5'
}

export function TetrateApiTccV1WorkflowsLBClassFromJSON(json: any): TetrateApiTccV1WorkflowsLBClass {
    return json as TetrateApiTccV1WorkflowsLBClass;
}

export function TetrateApiTccV1WorkflowsLBClassToJSON(value?: TetrateApiTccV1WorkflowsLBClass): any {
    return value as any;
}

