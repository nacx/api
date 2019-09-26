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
    TetrateApiTccCoreV1Service,
    TetrateApiTccCoreV1ServiceFromJSON,
    TetrateApiTccCoreV1ServiceToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1ListServicesResponse
 */
export interface TetrateApiTccCoreV1ListServicesResponse {
    /**
     * 
     * @type {Array<TetrateApiTccCoreV1Service>}
     * @memberof TetrateApiTccCoreV1ListServicesResponse
     */
    services?: Array<TetrateApiTccCoreV1Service>;
}

export function TetrateApiTccCoreV1ListServicesResponseFromJSON(json: any): TetrateApiTccCoreV1ListServicesResponse {
    return {
        'services': !exists(json, 'services') ? undefined : (json['services'] as Array<any>).map(TetrateApiTccCoreV1ServiceFromJSON),
    };
}

export function TetrateApiTccCoreV1ListServicesResponseToJSON(value?: TetrateApiTccCoreV1ListServicesResponse): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'services': value.services === undefined ? undefined : (value.services as Array<any>).map(TetrateApiTccCoreV1ServiceToJSON),
    };
}

