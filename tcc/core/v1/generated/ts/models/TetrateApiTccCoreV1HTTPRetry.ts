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
 * @interface TetrateApiTccCoreV1HTTPRetry
 */
export interface TetrateApiTccCoreV1HTTPRetry {
    /**
     * Number of retries for a given request. The interval between retries will be determined automatically (25ms+). Actual number of retries attempted depends on the httpReqTimeout.
     * @type {number}
     * @memberof TetrateApiTccCoreV1HTTPRetry
     */
    attempts?: number;
    /**
     * Timeout per retry attempt for a given request. format: 1h/1m/1s/1ms. MUST BE >=1ms.
     * @type {string}
     * @memberof TetrateApiTccCoreV1HTTPRetry
     */
    perTryTimeout?: string;
    /**
     * Specifies the conditions under which retry takes place. One or more policies can be specified using a ‘,’ delimited list. See the [supported policies](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#x-envoy-retry-on) and [here](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#x-envoy-retry-grpc-on) for more details.
     * @type {string}
     * @memberof TetrateApiTccCoreV1HTTPRetry
     */
    retryOn?: string;
}

export function TetrateApiTccCoreV1HTTPRetryFromJSON(json: any): TetrateApiTccCoreV1HTTPRetry {
    return {
        'attempts': !exists(json, 'attempts') ? undefined : json['attempts'],
        'perTryTimeout': !exists(json, 'perTryTimeout') ? undefined : json['perTryTimeout'],
        'retryOn': !exists(json, 'retryOn') ? undefined : json['retryOn'],
    };
}

export function TetrateApiTccCoreV1HTTPRetryToJSON(value?: TetrateApiTccCoreV1HTTPRetry): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'attempts': value.attempts,
        'perTryTimeout': value.perTryTimeout,
        'retryOn': value.retryOn,
    };
}


