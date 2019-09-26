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
    TetrateApiTccCoreV1ApplicationSpecificLB,
    TetrateApiTccCoreV1ApplicationSpecificLBFromJSON,
    TetrateApiTccCoreV1ApplicationSpecificLBToJSON,
    TetrateApiTccCoreV1ClientSettings,
    TetrateApiTccCoreV1ClientSettingsFromJSON,
    TetrateApiTccCoreV1ClientSettingsToJSON,
} from './';

/**
 * An Application is a collection of services. Each application typically corresponds to one or more kubernetes namespace or an application deployment on VMs.
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
     * short name for the application.
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
     * 
     * @type {TetrateApiTccCoreV1ClientSettings}
     * @memberof TetrateApiTccCoreV1Application
     */
    clientSettings?: TetrateApiTccCoreV1ClientSettings;
    /**
     * An application can be exposed on a shared load balancer in the cluster or a dedicated load balancer in one of the application\'s namespaces. Set app_lbs to indicate that the application should be exposed on the dedicated load balancers. Namespace where each of the dedicated load balancer is scoped should be unique.
     * @type {Array<TetrateApiTccCoreV1ApplicationSpecificLB>}
     * @memberof TetrateApiTccCoreV1Application
     */
    appLbs?: Array<TetrateApiTccCoreV1ApplicationSpecificLB>;
    /**
     * List of namespaces where the application services (or components) are scoped within. If omitted, the application is assumed to be scoped in a namespace matching the Id field.
     * @type {Array<string>}
     * @memberof TetrateApiTccCoreV1Application
     */
    namespaces?: Array<string>;
}

export function TetrateApiTccCoreV1ApplicationFromJSON(json: any): TetrateApiTccCoreV1Application {
    return {
        'name': !exists(json, 'name') ? undefined : json['name'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'clientSettings': !exists(json, 'clientSettings') ? undefined : TetrateApiTccCoreV1ClientSettingsFromJSON(json['clientSettings']),
        'appLbs': !exists(json, 'appLbs') ? undefined : (json['appLbs'] as Array<any>).map(TetrateApiTccCoreV1ApplicationSpecificLBFromJSON),
        'namespaces': !exists(json, 'namespaces') ? undefined : json['namespaces'],
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
        'clientSettings': TetrateApiTccCoreV1ClientSettingsToJSON(value.clientSettings),
        'appLbs': value.appLbs === undefined ? undefined : (value.appLbs as Array<any>).map(TetrateApiTccCoreV1ApplicationSpecificLBToJSON),
        'namespaces': value.namespaces,
    };
}


