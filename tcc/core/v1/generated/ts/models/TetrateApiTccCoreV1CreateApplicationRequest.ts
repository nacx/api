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
    TetrateApiTccCoreV1ClientSettings,
    TetrateApiTccCoreV1ClientSettingsFromJSON,
    TetrateApiTccCoreV1ClientSettingsToJSON,
    TetrateApiTccCoreV1CreateApplicationRequestApplicationSpecificLB,
    TetrateApiTccCoreV1CreateApplicationRequestApplicationSpecificLBFromJSON,
    TetrateApiTccCoreV1CreateApplicationRequestApplicationSpecificLBToJSON,
    TetrateApiTccCoreV1RoutingInfo,
    TetrateApiTccCoreV1RoutingInfoFromJSON,
    TetrateApiTccCoreV1RoutingInfoToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1CreateApplicationRequest
 */
export interface TetrateApiTccCoreV1CreateApplicationRequest {
    /**
     * Internal use only. Auto populated field.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateApplicationRequest
     */
    parent?: string;
    /**
     * Tenant.Id
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateApplicationRequest
     */
    tenant?: string;
    /**
     * Environment.Id
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateApplicationRequest
     */
    environment?: string;
    /**
     * short name for the application. Clusters are expected to have namespaces that match the application Id, especially on Kubernetes.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateApplicationRequest
     */
    id?: string;
    /**
     * Additional information.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateApplicationRequest
     */
    description?: string;
    /**
     * Hostname with which the application is exposed on a Gateway
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateApplicationRequest
     */
    hostname?: string;
    /**
     * 
     * @type {TetrateApiTccCoreV1RoutingInfo}
     * @memberof TetrateApiTccCoreV1CreateApplicationRequest
     */
    routingInfo?: TetrateApiTccCoreV1RoutingInfo;
    /**
     * 
     * @type {TetrateApiTccCoreV1ClientSettings}
     * @memberof TetrateApiTccCoreV1CreateApplicationRequest
     */
    clientSettings?: TetrateApiTccCoreV1ClientSettings;
    /**
     * 
     * @type {TetrateApiTccCoreV1CreateApplicationRequestApplicationSpecificLB}
     * @memberof TetrateApiTccCoreV1CreateApplicationRequest
     */
    appLb?: TetrateApiTccCoreV1CreateApplicationRequestApplicationSpecificLB;
}

export function TetrateApiTccCoreV1CreateApplicationRequestFromJSON(json: any): TetrateApiTccCoreV1CreateApplicationRequest {
    return {
        'parent': !exists(json, 'parent') ? undefined : json['parent'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'hostname': !exists(json, 'hostname') ? undefined : json['hostname'],
        'routingInfo': !exists(json, 'routingInfo') ? undefined : TetrateApiTccCoreV1RoutingInfoFromJSON(json['routingInfo']),
        'clientSettings': !exists(json, 'clientSettings') ? undefined : TetrateApiTccCoreV1ClientSettingsFromJSON(json['clientSettings']),
        'appLb': !exists(json, 'appLb') ? undefined : TetrateApiTccCoreV1CreateApplicationRequestApplicationSpecificLBFromJSON(json['appLb']),
    };
}

export function TetrateApiTccCoreV1CreateApplicationRequestToJSON(value?: TetrateApiTccCoreV1CreateApplicationRequest): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'parent': value.parent,
        'tenant': value.tenant,
        'environment': value.environment,
        'id': value.id,
        'description': value.description,
        'hostname': value.hostname,
        'routingInfo': TetrateApiTccCoreV1RoutingInfoToJSON(value.routingInfo),
        'clientSettings': TetrateApiTccCoreV1ClientSettingsToJSON(value.clientSettings),
        'appLb': TetrateApiTccCoreV1CreateApplicationRequestApplicationSpecificLBToJSON(value.appLb),
    };
}

