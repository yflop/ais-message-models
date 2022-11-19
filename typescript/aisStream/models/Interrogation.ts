/**
 * Ais-Stream WebsocketObjects
 * A sample API to illustrate OpenAPI concepts
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { InterrogationStation1Msg1 } from './InterrogationStation1Msg1';
import { InterrogationStation1Msg2 } from './InterrogationStation1Msg2';
import { InterrogationStation2 } from './InterrogationStation2';
import { HttpFile } from '../http/http';

export class Interrogation {
    'messageID': number;
    'repeatIndicator': number;
    'userID': number;
    'valid': boolean;
    'spare': number;
    'station1Msg1': InterrogationStation1Msg1;
    'station1Msg2': InterrogationStation1Msg2;
    'station2': InterrogationStation2;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "messageID",
            "baseName": "MessageID",
            "type": "number",
            "format": ""
        },
        {
            "name": "repeatIndicator",
            "baseName": "RepeatIndicator",
            "type": "number",
            "format": ""
        },
        {
            "name": "userID",
            "baseName": "UserID",
            "type": "number",
            "format": ""
        },
        {
            "name": "valid",
            "baseName": "Valid",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "spare",
            "baseName": "Spare",
            "type": "number",
            "format": ""
        },
        {
            "name": "station1Msg1",
            "baseName": "Station1Msg1",
            "type": "InterrogationStation1Msg1",
            "format": ""
        },
        {
            "name": "station1Msg2",
            "baseName": "Station1Msg2",
            "type": "InterrogationStation1Msg2",
            "format": ""
        },
        {
            "name": "station2",
            "baseName": "Station2",
            "type": "InterrogationStation2",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return Interrogation.attributeTypeMap;
    }

    public constructor() {
    }
}
