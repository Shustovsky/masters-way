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
import type { SchemasLabelInfo } from './SchemasLabelInfo';
import {
    SchemasLabelInfoFromJSON,
    SchemasLabelInfoFromJSONTyped,
    SchemasLabelInfoToJSON,
} from './SchemasLabelInfo';

/**
 * 
 * @export
 * @interface SchemasLabelStatistics
 */
export interface SchemasLabelStatistics {
    /**
     * 
     * @type {Array<SchemasLabelInfo>}
     * @memberof SchemasLabelStatistics
     */
    labels: Array<SchemasLabelInfo>;
}

/**
 * Check if a given object implements the SchemasLabelStatistics interface.
 */
export function instanceOfSchemasLabelStatistics(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "labels" in value;

    return isInstance;
}

export function SchemasLabelStatisticsFromJSON(json: any): SchemasLabelStatistics {
    return SchemasLabelStatisticsFromJSONTyped(json, false);
}

export function SchemasLabelStatisticsFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasLabelStatistics {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'labels': ((json['labels'] as Array<any>).map(SchemasLabelInfoFromJSON)),
    };
}


export function SchemasLabelStatisticsToJSON(value?: SchemasLabelStatistics | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'labels': ((value.labels as Array<any>).map(SchemasLabelInfoToJSON)),
    };
}

