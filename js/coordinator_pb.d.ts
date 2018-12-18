// package: 
// file: coordinator.proto

import * as jspb from "google-protobuf";

export class CoordinatorMessage extends jspb.Message {
  getType(): CoordinatorMessageType;
  setType(value: CoordinatorMessageType): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CoordinatorMessage.AsObject;
  static toObject(includeInstance: boolean, msg: CoordinatorMessage): CoordinatorMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CoordinatorMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CoordinatorMessage;
  static deserializeBinaryFromReader(message: CoordinatorMessage, reader: jspb.BinaryReader): CoordinatorMessage;
}

export namespace CoordinatorMessage {
  export type AsObject = {
    type: CoordinatorMessageType,
  }
}

export class WelcomeServerMessage extends jspb.Message {
  getType(): CoordinatorMessageType;
  setType(value: CoordinatorMessageType): void;

  getAlias(): string;
  setAlias(value: string): void;

  clearPeersList(): void;
  getPeersList(): Array<string>;
  setPeersList(value: Array<string>): void;
  addPeers(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WelcomeServerMessage.AsObject;
  static toObject(includeInstance: boolean, msg: WelcomeServerMessage): WelcomeServerMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WelcomeServerMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WelcomeServerMessage;
  static deserializeBinaryFromReader(message: WelcomeServerMessage, reader: jspb.BinaryReader): WelcomeServerMessage;
}

export namespace WelcomeServerMessage {
  export type AsObject = {
    type: CoordinatorMessageType,
    alias: string,
    peersList: Array<string>,
  }
}

export class WelcomeClientMessage extends jspb.Message {
  getType(): CoordinatorMessageType;
  setType(value: CoordinatorMessageType): void;

  getAlias(): string;
  setAlias(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WelcomeClientMessage.AsObject;
  static toObject(includeInstance: boolean, msg: WelcomeClientMessage): WelcomeClientMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WelcomeClientMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WelcomeClientMessage;
  static deserializeBinaryFromReader(message: WelcomeClientMessage, reader: jspb.BinaryReader): WelcomeClientMessage;
}

export namespace WelcomeClientMessage {
  export type AsObject = {
    type: CoordinatorMessageType,
    alias: string,
  }
}

export class ConnectMessage extends jspb.Message {
  getType(): CoordinatorMessageType;
  setType(value: CoordinatorMessageType): void;

  getToalias(): string;
  setToalias(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConnectMessage.AsObject;
  static toObject(includeInstance: boolean, msg: ConnectMessage): ConnectMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ConnectMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConnectMessage;
  static deserializeBinaryFromReader(message: ConnectMessage, reader: jspb.BinaryReader): ConnectMessage;
}

export namespace ConnectMessage {
  export type AsObject = {
    type: CoordinatorMessageType,
    toalias: string,
  }
}

export class WebRtcMessage extends jspb.Message {
  getType(): CoordinatorMessageType;
  setType(value: CoordinatorMessageType): void;

  getAlias(): string;
  setAlias(value: string): void;

  getToalias(): string;
  setToalias(value: string): void;

  getSdp(): string;
  setSdp(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WebRtcMessage.AsObject;
  static toObject(includeInstance: boolean, msg: WebRtcMessage): WebRtcMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WebRtcMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WebRtcMessage;
  static deserializeBinaryFromReader(message: WebRtcMessage, reader: jspb.BinaryReader): WebRtcMessage;
}

export namespace WebRtcMessage {
  export type AsObject = {
    type: CoordinatorMessageType,
    alias: string,
    toalias: string,
    sdp: string,
  }
}

export enum CoordinatorMessageType {
  UNKNOWN_MESSAGE_TYPE = 0,
  WELCOME_SERVER = 1,
  WELCOME_CLIENT = 2,
  CONNECT = 3,
  WEBRTC_SUPPORTED = 4,
  WEBRTC_OFFER = 5,
  WEBRTC_ANSWER = 6,
  WEBRTC_ICE_CANDIDATE = 7,
}

