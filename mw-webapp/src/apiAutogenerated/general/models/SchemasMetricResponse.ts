// @ts-nocheck
/* eslint-disable */
/**
 * Masters way general API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface SchemasMetricResponse
 */
export interface SchemasMetricResponse {
    /**
     * 
     * @type {string}
     * @memberof SchemasMetricResponse
     */
    description: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasMetricResponse
     */
    doneDate: string | null;
    /**
     * 
     * @type {number}
     * @memberof SchemasMetricResponse
     */
    estimationTime: number;
    /**
     * 
     * @type {boolean}
     * @memberof SchemasMetricResponse
     */
    isDone: boolean;
    /**
     * 
     * @type {string}
     * @memberof SchemasMetricResponse
     */
    uuid: string;
}

/**
 * Check if a given object implements the SchemasMetricResponse interface.
 */
export function instanceOfSchemasMetricResponse(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "description" in value;
    isInstance = isInstance && "doneDate" in value;
    isInstance = isInstance && "estimationTime" in value;
    isInstance = isInstance && "isDone" in value;
    isInstance = isInstance && "uuid" in value;

    return isInstance;
}

export function SchemasMetricResponseFromJSON(json: any): SchemasMetricResponse {
    return SchemasMetricResponseFromJSONTyped(json, false);
}

export function SchemasMetricResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasMetricResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'description': json['description'],
        'doneDate': json['doneDate'],
        'estimationTime': json['estimationTime'],
        'isDone': json['isDone'],
        'uuid': json['uuid'],
    };
}


export function SchemasMetricResponseToJSON(value?: SchemasMetricResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'description': value.description,
        'doneDate': value.doneDate,
        'estimationTime': value.estimationTime,
        'isDone': value.isDone,
        'uuid': value.uuid,
    };
}

