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
    TetrateApiTccCoreV1LoadBalancer,
    TetrateApiTccCoreV1LoadBalancerFromJSON,
    TetrateApiTccCoreV1LoadBalancerToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1ListLoadBalancersResponse
 */
export interface TetrateApiTccCoreV1ListLoadBalancersResponse {
    /**
     * 
     * @type {Array<TetrateApiTccCoreV1LoadBalancer>}
     * @memberof TetrateApiTccCoreV1ListLoadBalancersResponse
     */
    loadbalancers?: Array<TetrateApiTccCoreV1LoadBalancer>;
}

export function TetrateApiTccCoreV1ListLoadBalancersResponseFromJSON(json: any): TetrateApiTccCoreV1ListLoadBalancersResponse {
    return {
        'loadbalancers': !exists(json, 'loadbalancers') ? undefined : (json['loadbalancers'] as Array<any>).map(TetrateApiTccCoreV1LoadBalancerFromJSON),
    };
}

export function TetrateApiTccCoreV1ListLoadBalancersResponseToJSON(value?: TetrateApiTccCoreV1ListLoadBalancersResponse): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'loadbalancers': value.loadbalancers === undefined ? undefined : (value.loadbalancers as Array<any>).map(TetrateApiTccCoreV1LoadBalancerToJSON),
    };
}


