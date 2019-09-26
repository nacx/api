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
    TetrateApiTccCoreV1Cluster,
    TetrateApiTccCoreV1ClusterFromJSON,
    TetrateApiTccCoreV1ClusterToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1ListClustersResponse
 */
export interface TetrateApiTccCoreV1ListClustersResponse {
    /**
     * 
     * @type {Array<TetrateApiTccCoreV1Cluster>}
     * @memberof TetrateApiTccCoreV1ListClustersResponse
     */
    clusters?: Array<TetrateApiTccCoreV1Cluster>;
}

export function TetrateApiTccCoreV1ListClustersResponseFromJSON(json: any): TetrateApiTccCoreV1ListClustersResponse {
    return {
        'clusters': !exists(json, 'clusters') ? undefined : (json['clusters'] as Array<any>).map(TetrateApiTccCoreV1ClusterFromJSON),
    };
}

export function TetrateApiTccCoreV1ListClustersResponseToJSON(value?: TetrateApiTccCoreV1ListClustersResponse): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'clusters': value.clusters === undefined ? undefined : (value.clusters as Array<any>).map(TetrateApiTccCoreV1ClusterToJSON),
    };
}

