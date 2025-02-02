/* tslint:disable */
/* eslint-disable */
/**
 * Deepfence ThreatMapper
 * Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.
 *
 * The version of the OpenAPI document: 2.2.1
 * Contact: community@deepfence.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ModelCloudNodeAccountRegisterRespData } from './ModelCloudNodeAccountRegisterRespData';
import {
    ModelCloudNodeAccountRegisterRespDataFromJSON,
    ModelCloudNodeAccountRegisterRespDataFromJSONTyped,
    ModelCloudNodeAccountRegisterRespDataToJSON,
} from './ModelCloudNodeAccountRegisterRespData';

/**
 * 
 * @export
 * @interface ModelCloudNodeAccountRegisterResp
 */
export interface ModelCloudNodeAccountRegisterResp {
    /**
     * 
     * @type {ModelCloudNodeAccountRegisterRespData}
     * @memberof ModelCloudNodeAccountRegisterResp
     */
    data?: ModelCloudNodeAccountRegisterRespData;
}

/**
 * Check if a given object implements the ModelCloudNodeAccountRegisterResp interface.
 */
export function instanceOfModelCloudNodeAccountRegisterResp(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ModelCloudNodeAccountRegisterRespFromJSON(json: any): ModelCloudNodeAccountRegisterResp {
    return ModelCloudNodeAccountRegisterRespFromJSONTyped(json, false);
}

export function ModelCloudNodeAccountRegisterRespFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelCloudNodeAccountRegisterResp {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'data': !exists(json, 'data') ? undefined : ModelCloudNodeAccountRegisterRespDataFromJSON(json['data']),
    };
}

export function ModelCloudNodeAccountRegisterRespToJSON(value?: ModelCloudNodeAccountRegisterResp | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'data': ModelCloudNodeAccountRegisterRespDataToJSON(value.data),
    };
}

