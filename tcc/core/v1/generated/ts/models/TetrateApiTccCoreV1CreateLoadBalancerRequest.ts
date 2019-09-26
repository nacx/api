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
    TetrateApiTccCoreV1LoadBalancerClass,
    TetrateApiTccCoreV1LoadBalancerClassFromJSON,
    TetrateApiTccCoreV1LoadBalancerClassToJSON,
    TetrateApiTccCoreV1Registry,
    TetrateApiTccCoreV1RegistryFromJSON,
    TetrateApiTccCoreV1RegistryToJSON,
    TetrateApiTccCoreV1TLSSettings,
    TetrateApiTccCoreV1TLSSettingsFromJSON,
    TetrateApiTccCoreV1TLSSettingsToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1CreateLoadBalancerRequest
 */
export interface TetrateApiTccCoreV1CreateLoadBalancerRequest {
    /**
     * Internal use only. Auto populated field.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateLoadBalancerRequest
     */
    parent?: string;
    /**
     * Tenant.Id
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateLoadBalancerRequest
     */
    tenant?: string;
    /**
     * Environment.Id
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateLoadBalancerRequest
     */
    environment?: string;
    /**
     * if present, this will be used as the id for the created object.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateLoadBalancerRequest
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateLoadBalancerRequest
     */
    description?: string;
    /**
     * 
     * @type {boolean}
     * @memberof TetrateApiTccCoreV1CreateLoadBalancerRequest
     */
    enableWorkflows?: boolean;
    /**
     * 
     * @type {TetrateApiTccCoreV1LoadBalancerClass}
     * @memberof TetrateApiTccCoreV1CreateLoadBalancerRequest
     */
    loadBalancerClass?: TetrateApiTccCoreV1LoadBalancerClass;
    /**
     * 
     * @type {TetrateApiTccCoreV1Registry}
     * @memberof TetrateApiTccCoreV1CreateLoadBalancerRequest
     */
    registry?: TetrateApiTccCoreV1Registry;
    /**
     * The namespace where the load balancer is/will be deployed in a given cluster.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateLoadBalancerRequest
     */
    clusterNamespace?: string;
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof TetrateApiTccCoreV1CreateLoadBalancerRequest
     */
    labels?: { [key: string]: string; };
    /**
     * 
     * @type {{ [key: string]: TetrateApiTccCoreV1TLSSettings; }}
     * @memberof TetrateApiTccCoreV1CreateLoadBalancerRequest
     */
    applications?: { [key: string]: TetrateApiTccCoreV1TLSSettings; };
    /**
     * 
     * @type {TetrateApiTccCoreV1ClientSettings}
     * @memberof TetrateApiTccCoreV1CreateLoadBalancerRequest
     */
    clientSettings?: TetrateApiTccCoreV1ClientSettings;
}

export function TetrateApiTccCoreV1CreateLoadBalancerRequestFromJSON(json: any): TetrateApiTccCoreV1CreateLoadBalancerRequest {
    return {
        'parent': !exists(json, 'parent') ? undefined : json['parent'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'enableWorkflows': !exists(json, 'enableWorkflows') ? undefined : json['enableWorkflows'],
        'loadBalancerClass': !exists(json, 'loadBalancerClass') ? undefined : TetrateApiTccCoreV1LoadBalancerClassFromJSON(json['loadBalancerClass']),
        'registry': !exists(json, 'registry') ? undefined : TetrateApiTccCoreV1RegistryFromJSON(json['registry']),
        'clusterNamespace': !exists(json, 'clusterNamespace') ? undefined : json['clusterNamespace'],
        'labels': !exists(json, 'labels') ? undefined : json['labels'],
        'applications': !exists(json, 'applications') ? undefined : mapValues(json['applications'], TetrateApiTccCoreV1TLSSettingsFromJSON),
        'clientSettings': !exists(json, 'clientSettings') ? undefined : TetrateApiTccCoreV1ClientSettingsFromJSON(json['clientSettings']),
    };
}

export function TetrateApiTccCoreV1CreateLoadBalancerRequestToJSON(value?: TetrateApiTccCoreV1CreateLoadBalancerRequest): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'parent': value.parent,
        'tenant': value.tenant,
        'environment': value.environment,
        'id': value.id,
        'description': value.description,
        'enableWorkflows': value.enableWorkflows,
        'loadBalancerClass': TetrateApiTccCoreV1LoadBalancerClassToJSON(value.loadBalancerClass),
        'registry': TetrateApiTccCoreV1RegistryToJSON(value.registry),
        'clusterNamespace': value.clusterNamespace,
        'labels': value.labels,
        'applications': value.applications === undefined ? undefined : mapValues(value.applications, TetrateApiTccCoreV1TLSSettingsToJSON),
        'clientSettings': TetrateApiTccCoreV1ClientSettingsToJSON(value.clientSettings),
    };
}


