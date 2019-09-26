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
    TetrateApiTccCoreV1LBRouteSettingsLBRouteAuthSettings,
    TetrateApiTccCoreV1LBRouteSettingsLBRouteAuthSettingsFromJSON,
    TetrateApiTccCoreV1LBRouteSettingsLBRouteAuthSettingsToJSON,
    TetrateApiTccCoreV1RateLimitSettings,
    TetrateApiTccCoreV1RateLimitSettingsFromJSON,
    TetrateApiTccCoreV1RateLimitSettingsToJSON,
    TetrateApiTccCoreV1TLSSettings,
    TetrateApiTccCoreV1TLSSettingsFromJSON,
    TetrateApiTccCoreV1TLSSettingsToJSON,
    TetrateApiTccCoreV1TcpSettings,
    TetrateApiTccCoreV1TcpSettingsFromJSON,
    TetrateApiTccCoreV1TcpSettingsToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1LBRouteSettingsLBRoute
 */
export interface TetrateApiTccCoreV1LBRouteSettingsLBRoute {
    /**
     * Hostname with which the application is exposed on this load balancer.
     * @type {string}
     * @memberof TetrateApiTccCoreV1LBRouteSettingsLBRoute
     */
    hostname?: string;
    /**
     * 
     * @type {TetrateApiTccCoreV1TLSSettings}
     * @memberof TetrateApiTccCoreV1LBRouteSettingsLBRoute
     */
    tls?: TetrateApiTccCoreV1TLSSettings;
    /**
     * 
     * @type {TetrateApiTccCoreV1HttpSettings}
     * @memberof TetrateApiTccCoreV1LBRouteSettingsLBRoute
     */
    httpSettings?: TetrateApiTccCoreV1HttpSettings;
    /**
     * 
     * @type {TetrateApiTccCoreV1TcpSettings}
     * @memberof TetrateApiTccCoreV1LBRouteSettingsLBRoute
     */
    tcpSettings?: TetrateApiTccCoreV1TcpSettings;
    /**
     * One or more auth settings (OR-ed).
     * @type {Array<TetrateApiTccCoreV1LBRouteSettingsLBRouteAuthSettings>}
     * @memberof TetrateApiTccCoreV1LBRouteSettingsLBRoute
     */
    authSettings?: Array<TetrateApiTccCoreV1LBRouteSettingsLBRouteAuthSettings>;
    /**
     * 
     * @type {TetrateApiTccCoreV1RateLimitSettings}
     * @memberof TetrateApiTccCoreV1LBRouteSettingsLBRoute
     */
    rateLimitSettings?: TetrateApiTccCoreV1RateLimitSettings;
}

export function TetrateApiTccCoreV1LBRouteSettingsLBRouteFromJSON(json: any): TetrateApiTccCoreV1LBRouteSettingsLBRoute {
    return {
        'hostname': !exists(json, 'hostname') ? undefined : json['hostname'],
        'tls': !exists(json, 'tls') ? undefined : TetrateApiTccCoreV1TLSSettingsFromJSON(json['tls']),
        'httpSettings': !exists(json, 'httpSettings') ? undefined : TetrateApiTccCoreV1HttpSettingsFromJSON(json['httpSettings']),
        'tcpSettings': !exists(json, 'tcpSettings') ? undefined : TetrateApiTccCoreV1TcpSettingsFromJSON(json['tcpSettings']),
        'authSettings': !exists(json, 'authSettings') ? undefined : (json['authSettings'] as Array<any>).map(TetrateApiTccCoreV1LBRouteSettingsLBRouteAuthSettingsFromJSON),
        'rateLimitSettings': !exists(json, 'rateLimitSettings') ? undefined : TetrateApiTccCoreV1RateLimitSettingsFromJSON(json['rateLimitSettings']),
    };
}

export function TetrateApiTccCoreV1LBRouteSettingsLBRouteToJSON(value?: TetrateApiTccCoreV1LBRouteSettingsLBRoute): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'hostname': value.hostname,
        'tls': TetrateApiTccCoreV1TLSSettingsToJSON(value.tls),
        'httpSettings': TetrateApiTccCoreV1HttpSettingsToJSON(value.httpSettings),
        'tcpSettings': TetrateApiTccCoreV1TcpSettingsToJSON(value.tcpSettings),
        'authSettings': value.authSettings === undefined ? undefined : (value.authSettings as Array<any>).map(TetrateApiTccCoreV1LBRouteSettingsLBRouteAuthSettingsToJSON),
        'rateLimitSettings': TetrateApiTccCoreV1RateLimitSettingsToJSON(value.rateLimitSettings),
    };
}


