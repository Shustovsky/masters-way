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


import * as runtime from '../runtime';
import type {
  SchemasCompositeDayReportPopulatedResponse,
  SchemasCreateDayReportPayload,
  SchemasListDayReportsResponse,
} from '../models/index';
import {
    SchemasCompositeDayReportPopulatedResponseFromJSON,
    SchemasCompositeDayReportPopulatedResponseToJSON,
    SchemasCreateDayReportPayloadFromJSON,
    SchemasCreateDayReportPayloadToJSON,
    SchemasListDayReportsResponseFromJSON,
    SchemasListDayReportsResponseToJSON,
} from '../models/index';

export interface CreateDayReportRequest {
    request: SchemasCreateDayReportPayload;
}

export interface GetDayReportsRequest {
    wayId: string;
    page?: number;
    limit?: number;
}

/**
 * 
 */
export class DayReportApi extends runtime.BaseAPI {

    /**
     * Create a new dayReport
     */
    async createDayReportRaw(requestParameters: CreateDayReportRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasCompositeDayReportPopulatedResponse>> {
        if (requestParameters.request === null || requestParameters.request === undefined) {
            throw new runtime.RequiredError('request','Required parameter requestParameters.request was null or undefined when calling createDayReport.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/dayReports`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SchemasCreateDayReportPayloadToJSON(requestParameters.request),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasCompositeDayReportPopulatedResponseFromJSON(jsonValue));
    }

    /**
     * Create a new dayReport
     */
    async createDayReport(requestParameters: CreateDayReportRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasCompositeDayReportPopulatedResponse> {
        const response = await this.createDayReportRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get list of day reports by way UUID
     */
    async getDayReportsRaw(requestParameters: GetDayReportsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasListDayReportsResponse>> {
        if (requestParameters.wayId === null || requestParameters.wayId === undefined) {
            throw new runtime.RequiredError('wayId','Required parameter requestParameters.wayId was null or undefined when calling getDayReports.');
        }

        const queryParameters: any = {};

        if (requestParameters.page !== undefined) {
            queryParameters['page'] = requestParameters.page;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/dayReports/{wayId}`.replace(`{${"wayId"}}`, encodeURIComponent(String(requestParameters.wayId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasListDayReportsResponseFromJSON(jsonValue));
    }

    /**
     * Get list of day reports by way UUID
     */
    async getDayReports(requestParameters: GetDayReportsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasListDayReportsResponse> {
        const response = await this.getDayReportsRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
