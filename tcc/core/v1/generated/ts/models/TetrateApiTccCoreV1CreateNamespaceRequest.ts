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
 * @interface TetrateApiTccCoreV1CreateNamespaceRequest
 */
export interface TetrateApiTccCoreV1CreateNamespaceRequest {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateNamespaceRequest
     */
    parent?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateNamespaceRequest
     */
    tenant?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateNamespaceRequest
     */
    environment?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateNamespaceRequest
     */
    cluster?: string;
    /**
     * if present, this will be used as the id for the created object.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateNamespaceRequest
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateNamespaceRequest
     */
    description?: string;
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof TetrateApiTccCoreV1CreateNamespaceRequest
     */
    labels?: { [key: string]: string; };
}

export function TetrateApiTccCoreV1CreateNamespaceRequestFromJSON(json: any): TetrateApiTccCoreV1CreateNamespaceRequest {
    return {
        'parent': !exists(json, 'parent') ? undefined : json['parent'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'cluster': !exists(json, 'cluster') ? undefined : json['cluster'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'labels': !exists(json, 'labels') ? undefined : json['labels'],
    };
}

export function TetrateApiTccCoreV1CreateNamespaceRequestToJSON(value?: TetrateApiTccCoreV1CreateNamespaceRequest): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'parent': value.parent,
        'tenant': value.tenant,
        'environment': value.environment,
        'cluster': value.cluster,
        'id': value.id,
        'description': value.description,
        'labels': value.labels,
    };
}


