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
    TetrateApiTccCoreV1BulkLoadClusterRequestClusterWithNamespaces,
    TetrateApiTccCoreV1BulkLoadClusterRequestClusterWithNamespacesFromJSON,
    TetrateApiTccCoreV1BulkLoadClusterRequestClusterWithNamespacesToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1BulkLoadClusterRequest
 */
export interface TetrateApiTccCoreV1BulkLoadClusterRequest {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1BulkLoadClusterRequest
     */
    parent?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1BulkLoadClusterRequest
     */
    tenant?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1BulkLoadClusterRequest
     */
    environment?: string;
    /**
     * 
     * @type {Array<TetrateApiTccCoreV1BulkLoadClusterRequestClusterWithNamespaces>}
     * @memberof TetrateApiTccCoreV1BulkLoadClusterRequest
     */
    clusters?: Array<TetrateApiTccCoreV1BulkLoadClusterRequestClusterWithNamespaces>;
}

export function TetrateApiTccCoreV1BulkLoadClusterRequestFromJSON(json: any): TetrateApiTccCoreV1BulkLoadClusterRequest {
    return {
        'parent': !exists(json, 'parent') ? undefined : json['parent'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'clusters': !exists(json, 'clusters') ? undefined : (json['clusters'] as Array<any>).map(TetrateApiTccCoreV1BulkLoadClusterRequestClusterWithNamespacesFromJSON),
    };
}

export function TetrateApiTccCoreV1BulkLoadClusterRequestToJSON(value?: TetrateApiTccCoreV1BulkLoadClusterRequest): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'parent': value.parent,
        'tenant': value.tenant,
        'environment': value.environment,
        'clusters': value.clusters === undefined ? undefined : (value.clusters as Array<any>).map(TetrateApiTccCoreV1BulkLoadClusterRequestClusterWithNamespacesToJSON),
    };
}


