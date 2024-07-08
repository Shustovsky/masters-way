// @ts-nocheck
/* eslint-disable */
/**
 * Masters way API
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
 * @interface SchemasAIResponse
 */
export interface SchemasAIResponse {
    /**
     * 
     * @type {string}
     * @memberof SchemasAIResponse
     */
    answer?: string;
}

/**
 * Check if a given object implements the SchemasAIResponse interface.
 */
export function instanceOfSchemasAIResponse(
    value: object
): boolean {
    let isInstance = true;

    return isInstance;
}

export function SchemasAIResponseFromJSON(json: any): SchemasAIResponse {
    return SchemasAIResponseFromJSONTyped(json, false);
}

export function SchemasAIResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasAIResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'answer': !exists(json, 'answer') ? undefined : json['answer'],
    };
}


export function SchemasAIResponseToJSON(value?: SchemasAIResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'answer': value.answer,
    };
}
