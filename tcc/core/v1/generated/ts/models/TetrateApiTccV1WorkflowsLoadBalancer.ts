// tslint:disable
/**
 * configproducer.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists } from '../runtime';
import {
    TetrateApiTccV1WorkflowsLBClass,
    TetrateApiTccV1WorkflowsLBClassFromJSON,
    TetrateApiTccV1WorkflowsLBClassToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccV1WorkflowsLoadBalancer
 */
export interface TetrateApiTccV1WorkflowsLoadBalancer {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccV1WorkflowsLoadBalancer
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccV1WorkflowsLoadBalancer
     */
    workspace?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccV1WorkflowsLoadBalancer
     */
    tenant?: string;
    /**
     * 
     * @type {TetrateApiTccV1WorkflowsLBClass}
     * @memberof TetrateApiTccV1WorkflowsLoadBalancer
     */
    _class?: TetrateApiTccV1WorkflowsLBClass;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccV1WorkflowsLoadBalancer
     */
    managementIp?: string;
    /**
     * The name of the service that implements this load balancer behavior (i.e. the set of Envoys) e.g. tenant/123/workspace/456/service/ingress-envoys.foo.com.
     * @type {string}
     * @memberof TetrateApiTccV1WorkflowsLoadBalancer
     */
    serviceName?: string;
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof TetrateApiTccV1WorkflowsLoadBalancer
     */
    labels?: { [key: string]: string; };
    /**
     * 
     * @type {Array<string>}
     * @memberof TetrateApiTccV1WorkflowsLoadBalancer
     */
    services?: Array<string>;
}

export function TetrateApiTccV1WorkflowsLoadBalancerFromJSON(json: any): TetrateApiTccV1WorkflowsLoadBalancer {
    return {
        'name': !exists(json, 'name') ? undefined : json['name'],
        'workspace': !exists(json, 'workspace') ? undefined : json['workspace'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        '_class': !exists(json, 'class') ? undefined : TetrateApiTccV1WorkflowsLBClassFromJSON(json['class']),
        'managementIp': !exists(json, 'managementIp') ? undefined : json['managementIp'],
        'serviceName': !exists(json, 'serviceName') ? undefined : json['serviceName'],
        'labels': !exists(json, 'labels') ? undefined : json['labels'],
        'services': !exists(json, 'services') ? undefined : json['services'],
    };
}

export function TetrateApiTccV1WorkflowsLoadBalancerToJSON(value?: TetrateApiTccV1WorkflowsLoadBalancer): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'name': value.name,
        'workspace': value.workspace,
        'tenant': value.tenant,
        'class': TetrateApiTccV1WorkflowsLBClassToJSON(value._class),
        'managementIp': value.managementIp,
        'serviceName': value.serviceName,
        'labels': value.labels,
        'services': value.services,
    };
}


