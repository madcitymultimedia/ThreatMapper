/* tslint:disable */
/* eslint-disable */
/**
 * Deepfence ThreatMapper
 * Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.
 *
 * The version of the OpenAPI document: 2.0.0
 * Contact: community@deepfence.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ModelLoginResponse
 */
export interface ModelLoginResponse {
    /**
     * 
     * @type {string}
     * @memberof ModelLoginResponse
     */
    access_token: string;
    /**
     * 
     * @type {string}
     * @memberof ModelLoginResponse
     */
    email_domain: string;
    /**
     * 
     * @type {string}
     * @memberof ModelLoginResponse
     */
    license_key: string;
    /**
     * 
     * @type {boolean}
     * @memberof ModelLoginResponse
     */
    license_registered: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof ModelLoginResponse
     */
    onboarding_required: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof ModelLoginResponse
     */
    password_invalidated: boolean;
    /**
     * 
     * @type {string}
     * @memberof ModelLoginResponse
     */
    refresh_token: string;
}

/**
 * Check if a given object implements the ModelLoginResponse interface.
 */
export function instanceOfModelLoginResponse(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "access_token" in value;
    isInstance = isInstance && "email_domain" in value;
    isInstance = isInstance && "license_key" in value;
    isInstance = isInstance && "license_registered" in value;
    isInstance = isInstance && "onboarding_required" in value;
    isInstance = isInstance && "password_invalidated" in value;
    isInstance = isInstance && "refresh_token" in value;

    return isInstance;
}

export function ModelLoginResponseFromJSON(json: any): ModelLoginResponse {
    return ModelLoginResponseFromJSONTyped(json, false);
}

export function ModelLoginResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelLoginResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'access_token': json['access_token'],
        'email_domain': json['email_domain'],
        'license_key': json['license_key'],
        'license_registered': json['license_registered'],
        'onboarding_required': json['onboarding_required'],
        'password_invalidated': json['password_invalidated'],
        'refresh_token': json['refresh_token'],
    };
}

export function ModelLoginResponseToJSON(value?: ModelLoginResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'access_token': value.access_token,
        'email_domain': value.email_domain,
        'license_key': value.license_key,
        'license_registered': value.license_registered,
        'onboarding_required': value.onboarding_required,
        'password_invalidated': value.password_invalidated,
        'refresh_token': value.refresh_token,
    };
}

