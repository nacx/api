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
    TetrateApiTccCoreV1RouteDestination,
    TetrateApiTccCoreV1RouteDestinationFromJSON,
    TetrateApiTccCoreV1RouteDestinationToJSON,
} from './';

/**
 * One or more destinations in a local cluster for the given request.
 * @export
 * @interface TetrateApiTccCoreV1Route
 */
export interface TetrateApiTccCoreV1Route {
    /**
     * 
     * @type {Array<TetrateApiTccCoreV1RouteDestination>}
     * @memberof TetrateApiTccCoreV1Route
     */
    destinations?: Array<TetrateApiTccCoreV1RouteDestination>;
}

export function TetrateApiTccCoreV1RouteFromJSON(json: any): TetrateApiTccCoreV1Route {
    return {
        'destinations': !exists(json, 'destinations') ? undefined : (json['destinations'] as Array<any>).map(TetrateApiTccCoreV1RouteDestinationFromJSON),
    };
}

export function TetrateApiTccCoreV1RouteToJSON(value?: TetrateApiTccCoreV1Route): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'destinations': value.destinations === undefined ? undefined : (value.destinations as Array<any>).map(TetrateApiTccCoreV1RouteDestinationToJSON),
    };
}


