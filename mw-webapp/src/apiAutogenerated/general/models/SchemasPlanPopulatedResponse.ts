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
import type { SchemasJobTagResponse } from './SchemasJobTagResponse';
import {
    SchemasJobTagResponseFromJSON,
    SchemasJobTagResponseFromJSONTyped,
    SchemasJobTagResponseToJSON,
} from './SchemasJobTagResponse';

/**
 * 
 * @export
 * @interface SchemasPlanPopulatedResponse
 */
export interface SchemasPlanPopulatedResponse {
    /**
     * 
     * @type {string}
     * @memberof SchemasPlanPopulatedResponse
     */
    createdAt: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasPlanPopulatedResponse
     */
    dayReportUuid: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasPlanPopulatedResponse
     */
    description: string;
    /**
     * 
     * @type {boolean}
     * @memberof SchemasPlanPopulatedResponse
     */
    isDone: boolean;
    /**
     * 
     * @type {string}
     * @memberof SchemasPlanPopulatedResponse
     */
    ownerName: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasPlanPopulatedResponse
     */
    ownerUuid: string;
    /**
     * 
     * @type {Array<SchemasJobTagResponse>}
     * @memberof SchemasPlanPopulatedResponse
     */
    tags: Array<SchemasJobTagResponse>;
    /**
     * 
     * @type {number}
     * @memberof SchemasPlanPopulatedResponse
     */
    time: number;
    /**
     * 
     * @type {string}
     * @memberof SchemasPlanPopulatedResponse
     */
    updatedAt: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasPlanPopulatedResponse
     */
    uuid: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasPlanPopulatedResponse
     */
    wayName: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasPlanPopulatedResponse
     */
    wayUuid: string;
}

/**
 * Check if a given object implements the SchemasPlanPopulatedResponse interface.
 */
export function instanceOfSchemasPlanPopulatedResponse(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "createdAt" in value;
    isInstance = isInstance && "dayReportUuid" in value;
    isInstance = isInstance && "description" in value;
    isInstance = isInstance && "isDone" in value;
    isInstance = isInstance && "ownerName" in value;
    isInstance = isInstance && "ownerUuid" in value;
    isInstance = isInstance && "tags" in value;
    isInstance = isInstance && "time" in value;
    isInstance = isInstance && "updatedAt" in value;
    isInstance = isInstance && "uuid" in value;
    isInstance = isInstance && "wayName" in value;
    isInstance = isInstance && "wayUuid" in value;

    return isInstance;
}

export function SchemasPlanPopulatedResponseFromJSON(json: any): SchemasPlanPopulatedResponse {
    return SchemasPlanPopulatedResponseFromJSONTyped(json, false);
}

export function SchemasPlanPopulatedResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasPlanPopulatedResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'createdAt': json['createdAt'],
        'dayReportUuid': json['dayReportUuid'],
        'description': json['description'],
        'isDone': json['isDone'],
        'ownerName': json['ownerName'],
        'ownerUuid': json['ownerUuid'],
        'tags': ((json['tags'] as Array<any>).map(SchemasJobTagResponseFromJSON)),
        'time': json['time'],
        'updatedAt': json['updatedAt'],
        'uuid': json['uuid'],
        'wayName': json['wayName'],
        'wayUuid': json['wayUuid'],
    };
}


export function SchemasPlanPopulatedResponseToJSON(value?: SchemasPlanPopulatedResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'createdAt': value.createdAt,
        'dayReportUuid': value.dayReportUuid,
        'description': value.description,
        'isDone': value.isDone,
        'ownerName': value.ownerName,
        'ownerUuid': value.ownerUuid,
        'tags': ((value.tags as Array<any>).map(SchemasJobTagResponseToJSON)),
        'time': value.time,
        'updatedAt': value.updatedAt,
        'uuid': value.uuid,
        'wayName': value.wayName,
        'wayUuid': value.wayUuid,
    };
}

