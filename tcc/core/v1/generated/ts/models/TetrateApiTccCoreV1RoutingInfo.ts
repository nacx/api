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
    TetrateApiTccCoreV1HttpSettings,
    TetrateApiTccCoreV1HttpSettingsFromJSON,
    TetrateApiTccCoreV1HttpSettingsToJSON,
    TetrateApiTccCoreV1Subset,
    TetrateApiTccCoreV1SubsetFromJSON,
    TetrateApiTccCoreV1SubsetToJSON,
    TetrateApiTccCoreV1TcpSettings,
    TetrateApiTccCoreV1TcpSettingsFromJSON,
    TetrateApiTccCoreV1TcpSettingsToJSON,
} from './';

/**
 * HTTP routing settings for application or service. Subsets of a service should be declared here as well.
 * @export
 * @interface TetrateApiTccCoreV1RoutingInfo
 */
export interface TetrateApiTccCoreV1RoutingInfo {
    /**
     * One or more versions of the service. Each version has a distinct name and a set of labels that help uniquely identify the pods/VMs of that version.
     * @type {Array<TetrateApiTccCoreV1Subset>}
     * @memberof TetrateApiTccCoreV1RoutingInfo
     */
    subsets?: Array<TetrateApiTccCoreV1Subset>;
    /**
     * 
     * @type {TetrateApiTccCoreV1HttpSettings}
     * @memberof TetrateApiTccCoreV1RoutingInfo
     */
    httpSettings?: TetrateApiTccCoreV1HttpSettings;
    /**
     * 
     * @type {TetrateApiTccCoreV1TcpSettings}
     * @memberof TetrateApiTccCoreV1RoutingInfo
     */
    tcpSettings?: TetrateApiTccCoreV1TcpSettings;
}

export function TetrateApiTccCoreV1RoutingInfoFromJSON(json: any): TetrateApiTccCoreV1RoutingInfo {
    return {
        'subsets': !exists(json, 'subsets') ? undefined : (json['subsets'] as Array<any>).map(TetrateApiTccCoreV1SubsetFromJSON),
        'httpSettings': !exists(json, 'httpSettings') ? undefined : TetrateApiTccCoreV1HttpSettingsFromJSON(json['httpSettings']),
        'tcpSettings': !exists(json, 'tcpSettings') ? undefined : TetrateApiTccCoreV1TcpSettingsFromJSON(json['tcpSettings']),
    };
}

export function TetrateApiTccCoreV1RoutingInfoToJSON(value?: TetrateApiTccCoreV1RoutingInfo): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'subsets': value.subsets === undefined ? undefined : (value.subsets as Array<any>).map(TetrateApiTccCoreV1SubsetToJSON),
        'httpSettings': TetrateApiTccCoreV1HttpSettingsToJSON(value.httpSettings),
        'tcpSettings': TetrateApiTccCoreV1TcpSettingsToJSON(value.tcpSettings),
    };
}

