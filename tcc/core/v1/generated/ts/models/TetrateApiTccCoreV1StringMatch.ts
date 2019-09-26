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
/**
 * Describes how to match a given string in HTTP headers. Match is case-sensitive.
 * @export
 * @interface TetrateApiTccCoreV1StringMatch
 */
export interface TetrateApiTccCoreV1StringMatch {
    /**
     * exact string match
     * @type {string}
     * @memberof TetrateApiTccCoreV1StringMatch
     */
    exact?: string;
    /**
     * prefix-based match
     * @type {string}
     * @memberof TetrateApiTccCoreV1StringMatch
     */
    prefix?: string;
    /**
     * ECMAscript style regex-based match
     * @type {string}
     * @memberof TetrateApiTccCoreV1StringMatch
     */
    regex?: string;
}

export function TetrateApiTccCoreV1StringMatchFromJSON(json: any): TetrateApiTccCoreV1StringMatch {
    return {
        'exact': !exists(json, 'exact') ? undefined : json['exact'],
        'prefix': !exists(json, 'prefix') ? undefined : json['prefix'],
        'regex': !exists(json, 'regex') ? undefined : json['regex'],
    };
}

export function TetrateApiTccCoreV1StringMatchToJSON(value?: TetrateApiTccCoreV1StringMatch): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'exact': value.exact,
        'prefix': value.prefix,
        'regex': value.regex,
    };
}


