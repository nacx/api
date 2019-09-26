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
    TetrateApiTccCoreV1ClientSettingsSensitivity,
    TetrateApiTccCoreV1ClientSettingsSensitivityFromJSON,
    TetrateApiTccCoreV1ClientSettingsSensitivityToJSON,
    TetrateApiTccCoreV1HTTPRetry,
    TetrateApiTccCoreV1HTTPRetryFromJSON,
    TetrateApiTccCoreV1HTTPRetryToJSON,
} from './';

/**
 * All the timeouts, retries, circuit breakers etc. that consumers want to set to gaurd against failures of their dependencies. These roughly translate to pieces of istio virtual service, destination rules, etc. At runtime, for a given namespace, for the namespace\'s dependencies, we combine the client specified reliability settings with the virtual service/dest rule for the dependencies to produce an updated virtual service/dest rule per dependent service. All these client customized virtual services/dest rules will be private to the namespace i.e. will have an exportTo setting \'.\' . Doing so will make Pilot apply these customizations only to the namespace concerned without leaking it to all other namespaces.
 * @export
 * @interface TetrateApiTccCoreV1ClientSettings
 */
export interface TetrateApiTccCoreV1ClientSettings {
    /**
     * Timeout for HTTP requests.
     * @type {string}
     * @memberof TetrateApiTccCoreV1ClientSettings
     */
    httpRequestTimeout?: string;
    /**
     * 
     * @type {TetrateApiTccCoreV1HTTPRetry}
     * @memberof TetrateApiTccCoreV1ClientSettings
     */
    httpRetries?: TetrateApiTccCoreV1HTTPRetry;
    /**
     * TCP connection timeout.
     * @type {string}
     * @memberof TetrateApiTccCoreV1ClientSettings
     */
    tcpConnectTimeout?: string;
    /**
     * If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
     * @type {boolean}
     * @memberof TetrateApiTccCoreV1ClientSettings
     */
    tcpKeepalive?: boolean;
    /**
     * 
     * @type {TetrateApiTccCoreV1ClientSettingsSensitivity}
     * @memberof TetrateApiTccCoreV1ClientSettings
     */
    circuitBreakerSensitivity?: TetrateApiTccCoreV1ClientSettingsSensitivity;
}

export function TetrateApiTccCoreV1ClientSettingsFromJSON(json: any): TetrateApiTccCoreV1ClientSettings {
    return {
        'httpRequestTimeout': !exists(json, 'httpRequestTimeout') ? undefined : json['httpRequestTimeout'],
        'httpRetries': !exists(json, 'httpRetries') ? undefined : TetrateApiTccCoreV1HTTPRetryFromJSON(json['httpRetries']),
        'tcpConnectTimeout': !exists(json, 'tcpConnectTimeout') ? undefined : json['tcpConnectTimeout'],
        'tcpKeepalive': !exists(json, 'tcpKeepalive') ? undefined : json['tcpKeepalive'],
        'circuitBreakerSensitivity': !exists(json, 'circuitBreakerSensitivity') ? undefined : TetrateApiTccCoreV1ClientSettingsSensitivityFromJSON(json['circuitBreakerSensitivity']),
    };
}

export function TetrateApiTccCoreV1ClientSettingsToJSON(value?: TetrateApiTccCoreV1ClientSettings): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'httpRequestTimeout': value.httpRequestTimeout,
        'httpRetries': TetrateApiTccCoreV1HTTPRetryToJSON(value.httpRetries),
        'tcpConnectTimeout': value.tcpConnectTimeout,
        'tcpKeepalive': value.tcpKeepalive,
        'circuitBreakerSensitivity': TetrateApiTccCoreV1ClientSettingsSensitivityToJSON(value.circuitBreakerSensitivity),
    };
}


