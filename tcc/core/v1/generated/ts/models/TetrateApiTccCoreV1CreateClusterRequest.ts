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
    TetrateApiTccCoreV1Registry,
    TetrateApiTccCoreV1RegistryFromJSON,
    TetrateApiTccCoreV1RegistryToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1CreateClusterRequest
 */
export interface TetrateApiTccCoreV1CreateClusterRequest {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateClusterRequest
     */
    parent?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateClusterRequest
     */
    tenant?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateClusterRequest
     */
    environment?: string;
    /**
     * if present, this will be used as the id for the created object.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateClusterRequest
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateClusterRequest
     */
    description?: string;
    /**
     * 
     * @type {TetrateApiTccCoreV1Registry}
     * @memberof TetrateApiTccCoreV1CreateClusterRequest
     */
    registry?: TetrateApiTccCoreV1Registry;
    /**
     * Information like datacenter where the cluster is present.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateClusterRequest
     */
    country?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateClusterRequest
     */
    datacenter?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateClusterRequest
     */
    availabilityZone?: string;
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof TetrateApiTccCoreV1CreateClusterRequest
     */
    labels?: { [key: string]: string; };
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateClusterRequest
     */
    kubernetesClusterDomain?: string;
}

export function TetrateApiTccCoreV1CreateClusterRequestFromJSON(json: any): TetrateApiTccCoreV1CreateClusterRequest {
    return {
        'parent': !exists(json, 'parent') ? undefined : json['parent'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'registry': !exists(json, 'registry') ? undefined : TetrateApiTccCoreV1RegistryFromJSON(json['registry']),
        'country': !exists(json, 'country') ? undefined : json['country'],
        'datacenter': !exists(json, 'datacenter') ? undefined : json['datacenter'],
        'availabilityZone': !exists(json, 'availabilityZone') ? undefined : json['availabilityZone'],
        'labels': !exists(json, 'labels') ? undefined : json['labels'],
        'kubernetesClusterDomain': !exists(json, 'kubernetesClusterDomain') ? undefined : json['kubernetesClusterDomain'],
    };
}

export function TetrateApiTccCoreV1CreateClusterRequestToJSON(value?: TetrateApiTccCoreV1CreateClusterRequest): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'parent': value.parent,
        'tenant': value.tenant,
        'environment': value.environment,
        'id': value.id,
        'description': value.description,
        'registry': TetrateApiTccCoreV1RegistryToJSON(value.registry),
        'country': value.country,
        'datacenter': value.datacenter,
        'availabilityZone': value.availabilityZone,
        'labels': value.labels,
        'kubernetesClusterDomain': value.kubernetesClusterDomain,
    };
}


