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
    TetrateApiTccCoreV1TLSSettings,
    TetrateApiTccCoreV1TLSSettingsFromJSON,
    TetrateApiTccCoreV1TLSSettingsToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1ApplicationApplicationSpecificLB
 */
export interface TetrateApiTccCoreV1ApplicationApplicationSpecificLB {
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof TetrateApiTccCoreV1ApplicationApplicationSpecificLB
     */
    labels?: { [key: string]: string; };
    /**
     * 
     * @type {TetrateApiTccCoreV1TLSSettings}
     * @memberof TetrateApiTccCoreV1ApplicationApplicationSpecificLB
     */
    tls?: TetrateApiTccCoreV1TLSSettings;
}

export function TetrateApiTccCoreV1ApplicationApplicationSpecificLBFromJSON(json: any): TetrateApiTccCoreV1ApplicationApplicationSpecificLB {
    return {
        'labels': !exists(json, 'labels') ? undefined : json['labels'],
        'tls': !exists(json, 'tls') ? undefined : TetrateApiTccCoreV1TLSSettingsFromJSON(json['tls']),
    };
}

export function TetrateApiTccCoreV1ApplicationApplicationSpecificLBToJSON(value?: TetrateApiTccCoreV1ApplicationApplicationSpecificLB): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'labels': value.labels,
        'tls': TetrateApiTccCoreV1TLSSettingsToJSON(value.tls),
    };
}

