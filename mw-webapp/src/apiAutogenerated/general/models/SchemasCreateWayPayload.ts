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
 * @interface SchemasCreateWayPayload
 */
export interface SchemasCreateWayPayload {
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateWayPayload
     */
    copiedFromWayUuid: string | null;
    /**
     * 
     * @type {number}
     * @memberof SchemasCreateWayPayload
     */
    estimationTime: number;
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateWayPayload
     */
    goalDescription: string;
    /**
     * 
     * @type {boolean}
     * @memberof SchemasCreateWayPayload
     */
    isCompleted: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof SchemasCreateWayPayload
     */
    isPrivate: boolean;
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateWayPayload
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateWayPayload
     */
    ownerUuid: string;
}

/**
 * Check if a given object implements the SchemasCreateWayPayload interface.
 */
export function instanceOfSchemasCreateWayPayload(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "copiedFromWayUuid" in value;
    isInstance = isInstance && "estimationTime" in value;
    isInstance = isInstance && "goalDescription" in value;
    isInstance = isInstance && "isCompleted" in value;
    isInstance = isInstance && "isPrivate" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "ownerUuid" in value;

    return isInstance;
}

export function SchemasCreateWayPayloadFromJSON(json: any): SchemasCreateWayPayload {
    return SchemasCreateWayPayloadFromJSONTyped(json, false);
}

export function SchemasCreateWayPayloadFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasCreateWayPayload {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'copiedFromWayUuid': json['copiedFromWayUuid'],
        'estimationTime': json['estimationTime'],
        'goalDescription': json['goalDescription'],
        'isCompleted': json['isCompleted'],
        'isPrivate': json['isPrivate'],
        'name': json['name'],
        'ownerUuid': json['ownerUuid'],
    };
}


export function SchemasCreateWayPayloadToJSON(value?: SchemasCreateWayPayload | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'copiedFromWayUuid': value.copiedFromWayUuid,
        'estimationTime': value.estimationTime,
        'goalDescription': value.goalDescription,
        'isCompleted': value.isCompleted,
        'isPrivate': value.isPrivate,
        'name': value.name,
        'ownerUuid': value.ownerUuid,
    };
}

