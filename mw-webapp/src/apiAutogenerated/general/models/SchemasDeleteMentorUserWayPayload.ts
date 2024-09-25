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
 * @interface SchemasDeleteMentorUserWayPayload
 */
export interface SchemasDeleteMentorUserWayPayload {
    /**
     * 
     * @type {string}
     * @memberof SchemasDeleteMentorUserWayPayload
     */
    userUuid: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasDeleteMentorUserWayPayload
     */
    wayUuid: string;
}

/**
 * Check if a given object implements the SchemasDeleteMentorUserWayPayload interface.
 */
export function instanceOfSchemasDeleteMentorUserWayPayload(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "userUuid" in value;
    isInstance = isInstance && "wayUuid" in value;

    return isInstance;
}

export function SchemasDeleteMentorUserWayPayloadFromJSON(json: any): SchemasDeleteMentorUserWayPayload {
    return SchemasDeleteMentorUserWayPayloadFromJSONTyped(json, false);
}

export function SchemasDeleteMentorUserWayPayloadFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasDeleteMentorUserWayPayload {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'userUuid': json['userUuid'],
        'wayUuid': json['wayUuid'],
    };
}


export function SchemasDeleteMentorUserWayPayloadToJSON(value?: SchemasDeleteMentorUserWayPayload | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'userUuid': value.userUuid,
        'wayUuid': value.wayUuid,
    };
}

