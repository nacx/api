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
    TetrateApiTccCoreV1ApplicationApplicationSpecificLB,
    TetrateApiTccCoreV1ApplicationApplicationSpecificLBFromJSON,
    TetrateApiTccCoreV1ApplicationApplicationSpecificLBToJSON,
    TetrateApiTccCoreV1ClientSettings,
    TetrateApiTccCoreV1ClientSettingsFromJSON,
    TetrateApiTccCoreV1ClientSettingsToJSON,
    TetrateApiTccCoreV1RoutingInfo,
    TetrateApiTccCoreV1RoutingInfoFromJSON,
    TetrateApiTccCoreV1RoutingInfoToJSON,
} from './';

/**
 * An Application is a collection of services. Each application corresponds to a kubernetes namespace or an application deployment on VMs. The hostname of the application is exposed via the load balancer (either shared or dedicated).
 * @export
 * @interface TetrateApiTccCoreV1Application
 */
export interface TetrateApiTccCoreV1Application {
    /**
     * Internal use only. Auto populated field.
     * @type {string}
     * @memberof TetrateApiTccCoreV1Application
     */
    name?: string;
    /**
     * Tenant.Id.
     * @type {string}
     * @memberof TetrateApiTccCoreV1Application
     */
    tenant?: string;
    /**
     * Environment.Id.
     * @type {string}
     * @memberof TetrateApiTccCoreV1Application
     */
    environment?: string;
    /**
     * short name for the application. Clusters are expected to have namespaces that match the application Id, especially on Kubernetes.
     * @type {string}
     * @memberof TetrateApiTccCoreV1Application
     */
    id?: string;
    /**
     * Additional information.
     * @type {string}
     * @memberof TetrateApiTccCoreV1Application
     */
    description?: string;
    /**
     * Hostname with which the application is exposed on a Gateway.
     * @type {string}
     * @memberof TetrateApiTccCoreV1Application
     */
    hostname?: string;
    /**
     * 
     * @type {TetrateApiTccCoreV1RoutingInfo}
     * @memberof TetrateApiTccCoreV1Application
     */
    routingInfo?: TetrateApiTccCoreV1RoutingInfo;
    /**
     * 
     * @type {TetrateApiTccCoreV1ClientSettings}
     * @memberof TetrateApiTccCoreV1Application
     */
    clientSettings?: TetrateApiTccCoreV1ClientSettings;
    /**
     * 
     * @type {TetrateApiTccCoreV1ApplicationApplicationSpecificLB}
     * @memberof TetrateApiTccCoreV1Application
     */
    appLb?: TetrateApiTccCoreV1ApplicationApplicationSpecificLB;
}

export function TetrateApiTccCoreV1ApplicationFromJSON(json: any): TetrateApiTccCoreV1Application {
    return {
        'name': !exists(json, 'name') ? undefined : json['name'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'hostname': !exists(json, 'hostname') ? undefined : json['hostname'],
        'routingInfo': !exists(json, 'routingInfo') ? undefined : TetrateApiTccCoreV1RoutingInfoFromJSON(json['routingInfo']),
        'clientSettings': !exists(json, 'clientSettings') ? undefined : TetrateApiTccCoreV1ClientSettingsFromJSON(json['clientSettings']),
        'appLb': !exists(json, 'appLb') ? undefined : TetrateApiTccCoreV1ApplicationApplicationSpecificLBFromJSON(json['appLb']),
    };
}

export function TetrateApiTccCoreV1ApplicationToJSON(value?: TetrateApiTccCoreV1Application): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'name': value.name,
        'tenant': value.tenant,
        'environment': value.environment,
        'id': value.id,
        'description': value.description,
        'hostname': value.hostname,
        'routingInfo': TetrateApiTccCoreV1RoutingInfoToJSON(value.routingInfo),
        'clientSettings': TetrateApiTccCoreV1ClientSettingsToJSON(value.clientSettings),
        'appLb': TetrateApiTccCoreV1ApplicationApplicationSpecificLBToJSON(value.appLb),
    };
}


