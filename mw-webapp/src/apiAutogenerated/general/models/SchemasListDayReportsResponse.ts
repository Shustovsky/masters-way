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
import type { SchemasCompositeDayReportPopulatedResponse } from './SchemasCompositeDayReportPopulatedResponse';
import {
    SchemasCompositeDayReportPopulatedResponseFromJSON,
    SchemasCompositeDayReportPopulatedResponseFromJSONTyped,
    SchemasCompositeDayReportPopulatedResponseToJSON,
} from './SchemasCompositeDayReportPopulatedResponse';

/**
 * 
 * @export
 * @interface SchemasListDayReportsResponse
 */
export interface SchemasListDayReportsResponse {
    /**
     * 
     * @type {Array<SchemasCompositeDayReportPopulatedResponse>}
     * @memberof SchemasListDayReportsResponse
     */
    dayReports: Array<SchemasCompositeDayReportPopulatedResponse>;
    /**
     * 
     * @type {number}
     * @memberof SchemasListDayReportsResponse
     */
    size: number;
}

/**
 * Check if a given object implements the SchemasListDayReportsResponse interface.
 */
export function instanceOfSchemasListDayReportsResponse(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "dayReports" in value;
    isInstance = isInstance && "size" in value;

    return isInstance;
}

export function SchemasListDayReportsResponseFromJSON(json: any): SchemasListDayReportsResponse {
    return SchemasListDayReportsResponseFromJSONTyped(json, false);
}

export function SchemasListDayReportsResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasListDayReportsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dayReports': ((json['dayReports'] as Array<any>).map(SchemasCompositeDayReportPopulatedResponseFromJSON)),
        'size': json['size'],
    };
}


export function SchemasListDayReportsResponseToJSON(value?: SchemasListDayReportsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dayReports': ((value.dayReports as Array<any>).map(SchemasCompositeDayReportPopulatedResponseToJSON)),
        'size': value.size,
    };
}

