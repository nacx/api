// tslint:disable
/**
 * TCC Core API
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
    TetrateApiTccCoreV1HttpSettingsHTTPCookie,
    TetrateApiTccCoreV1HttpSettingsHTTPCookieFromJSON,
    TetrateApiTccCoreV1HttpSettingsHTTPCookieToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1HttpSettingsStickySession
 */
export interface TetrateApiTccCoreV1HttpSettingsStickySession {
    /**
     * Hash based on a specific HTTP header.
     * @type {string}
     * @memberof TetrateApiTccCoreV1HttpSettingsStickySession
     */
    header?: string;
    /**
     * 
     * @type {TetrateApiTccCoreV1HttpSettingsHTTPCookie}
     * @memberof TetrateApiTccCoreV1HttpSettingsStickySession
     */
    cookie?: TetrateApiTccCoreV1HttpSettingsHTTPCookie;
    /**
     * Hash based on the source IP address.
     * @type {boolean}
     * @memberof TetrateApiTccCoreV1HttpSettingsStickySession
     */
    useSourceIp?: boolean;
}

export function TetrateApiTccCoreV1HttpSettingsStickySessionFromJSON(json: any): TetrateApiTccCoreV1HttpSettingsStickySession {
    return {
        'header': !exists(json, 'header') ? undefined : json['header'],
        'cookie': !exists(json, 'cookie') ? undefined : TetrateApiTccCoreV1HttpSettingsHTTPCookieFromJSON(json['cookie']),
        'useSourceIp': !exists(json, 'useSourceIp') ? undefined : json['useSourceIp'],
    };
}

export function TetrateApiTccCoreV1HttpSettingsStickySessionToJSON(value?: TetrateApiTccCoreV1HttpSettingsStickySession): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'header': value.header,
        'cookie': TetrateApiTccCoreV1HttpSettingsHTTPCookieToJSON(value.cookie),
        'useSourceIp': value.useSourceIp,
    };
}


