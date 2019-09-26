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
/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1CorsPolicy
 */
export interface TetrateApiTccCoreV1CorsPolicy {
    /**
     * The list of origins that are allowed to perform CORS requests. The content will be serialized into the Access-Control-Allow-Origin header. Wildcard * will allow all origins.
     * @type {Array<string>}
     * @memberof TetrateApiTccCoreV1CorsPolicy
     */
    allowOrigin?: Array<string>;
    /**
     * List of HTTP methods allowed to access the resource. The content will be serialized into the Access-Control-Allow-Methods header.
     * @type {Array<string>}
     * @memberof TetrateApiTccCoreV1CorsPolicy
     */
    allowMethods?: Array<string>;
    /**
     * List of HTTP headers that can be used when requesting the resource. Serialized to Access-Control-Allow-Headers header.
     * @type {Array<string>}
     * @memberof TetrateApiTccCoreV1CorsPolicy
     */
    allowHeaders?: Array<string>;
    /**
     * A white list of HTTP headers that the browsers are allowed to access. Serialized into Access-Control-Expose-Headers header.
     * @type {Array<string>}
     * @memberof TetrateApiTccCoreV1CorsPolicy
     */
    exposeHeaders?: Array<string>;
    /**
     * Specifies how long the results of a preflight request can be cached. Translates to the Access-Control-Max-Age header.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CorsPolicy
     */
    maxAge?: string;
    /**
     * Indicates whether the caller is allowed to send the actual request (not the preflight) using credentials. Translates to Access-Control-Allow-Credentials header.
     * @type {boolean}
     * @memberof TetrateApiTccCoreV1CorsPolicy
     */
    allowCredentials?: boolean;
}

export function TetrateApiTccCoreV1CorsPolicyFromJSON(json: any): TetrateApiTccCoreV1CorsPolicy {
    return {
        'allowOrigin': !exists(json, 'allowOrigin') ? undefined : json['allowOrigin'],
        'allowMethods': !exists(json, 'allowMethods') ? undefined : json['allowMethods'],
        'allowHeaders': !exists(json, 'allowHeaders') ? undefined : json['allowHeaders'],
        'exposeHeaders': !exists(json, 'exposeHeaders') ? undefined : json['exposeHeaders'],
        'maxAge': !exists(json, 'maxAge') ? undefined : json['maxAge'],
        'allowCredentials': !exists(json, 'allowCredentials') ? undefined : json['allowCredentials'],
    };
}

export function TetrateApiTccCoreV1CorsPolicyToJSON(value?: TetrateApiTccCoreV1CorsPolicy): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'allowOrigin': value.allowOrigin,
        'allowMethods': value.allowMethods,
        'allowHeaders': value.allowHeaders,
        'exposeHeaders': value.exposeHeaders,
        'maxAge': value.maxAge,
        'allowCredentials': value.allowCredentials,
    };
}


