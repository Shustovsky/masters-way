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
 * @interface SchemasCreatePlanPayload
 */
export interface SchemasCreatePlanPayload {
    /**
     * 
     * @type {string}
     * @memberof SchemasCreatePlanPayload
     */
    dayReportUuid: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasCreatePlanPayload
     */
    description: string;
    /**
     * 
     * @type {boolean}
     * @memberof SchemasCreatePlanPayload
     */
    isDone: boolean;
    /**
     * 
     * @type {string}
     * @memberof SchemasCreatePlanPayload
     */
    ownerUuid: string;
    /**
     * 
     * @type {number}
     * @memberof SchemasCreatePlanPayload
     */
    time: number;
}

/**
 * Check if a given object implements the SchemasCreatePlanPayload interface.
 */
export function instanceOfSchemasCreatePlanPayload(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "dayReportUuid" in value;
    isInstance = isInstance && "description" in value;
    isInstance = isInstance && "isDone" in value;
    isInstance = isInstance && "ownerUuid" in value;
    isInstance = isInstance && "time" in value;

    return isInstance;
}

export function SchemasCreatePlanPayloadFromJSON(json: any): SchemasCreatePlanPayload {
    return SchemasCreatePlanPayloadFromJSONTyped(json, false);
}

export function SchemasCreatePlanPayloadFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasCreatePlanPayload {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dayReportUuid': json['dayReportUuid'],
        'description': json['description'],
        'isDone': json['isDone'],
        'ownerUuid': json['ownerUuid'],
        'time': json['time'],
    };
}


export function SchemasCreatePlanPayloadToJSON(value?: SchemasCreatePlanPayload | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dayReportUuid': value.dayReportUuid,
        'description': value.description,
        'isDone': value.isDone,
        'ownerUuid': value.ownerUuid,
        'time': value.time,
    };
}

