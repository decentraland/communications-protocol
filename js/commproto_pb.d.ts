// package: 
// file: commproto.proto

import * as jspb from "google-protobuf";

export class CoordinatorMessage extends jspb.Message {
  getType(): MessageType;
  setType(value: MessageType): void;

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
    type: MessageType,
  }
}

export class WelcomeServerMessage extends jspb.Message {
  getType(): MessageType;
  setType(value: MessageType): void;

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
    type: MessageType,
    alias: string,
    peersList: Array<string>,
  }
}

export class WelcomeClientMessage extends jspb.Message {
  getType(): MessageType;
  setType(value: MessageType): void;

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
    type: MessageType,
    alias: string,
  }
}

export class ConnectMessage extends jspb.Message {
  getType(): MessageType;
  setType(value: MessageType): void;

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
    type: MessageType,
    toalias: string,
  }
}

export class WebRtcMessage extends jspb.Message {
  getType(): MessageType;
  setType(value: MessageType): void;

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
    type: MessageType,
    alias: string,
    toalias: string,
    sdp: string,
  }
}

export class WorldCommMessage extends jspb.Message {
  getType(): MessageType;
  setType(value: MessageType): void;

  getTime(): number;
  setTime(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WorldCommMessage.AsObject;
  static toObject(includeInstance: boolean, msg: WorldCommMessage): WorldCommMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WorldCommMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WorldCommMessage;
  static deserializeBinaryFromReader(message: WorldCommMessage, reader: jspb.BinaryReader): WorldCommMessage;
}

export namespace WorldCommMessage {
  export type AsObject = {
    type: MessageType,
    time: number,
  }
}

export class PingMessage extends jspb.Message {
  getType(): MessageType;
  setType(value: MessageType): void;

  getTime(): number;
  setTime(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingMessage.AsObject;
  static toObject(includeInstance: boolean, msg: PingMessage): PingMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PingMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingMessage;
  static deserializeBinaryFromReader(message: PingMessage, reader: jspb.BinaryReader): PingMessage;
}

export namespace PingMessage {
  export type AsObject = {
    type: MessageType,
    time: number,
  }
}

export class ChangeTopicMessage extends jspb.Message {
  getType(): MessageType;
  setType(value: MessageType): void;

  getTime(): number;
  setTime(value: number): void;

  getTopic(): string;
  setTopic(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChangeTopicMessage.AsObject;
  static toObject(includeInstance: boolean, msg: ChangeTopicMessage): ChangeTopicMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ChangeTopicMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChangeTopicMessage;
  static deserializeBinaryFromReader(message: ChangeTopicMessage, reader: jspb.BinaryReader): ChangeTopicMessage;
}

export namespace ChangeTopicMessage {
  export type AsObject = {
    type: MessageType,
    time: number,
    topic: string,
  }
}

export class TopicMessage extends jspb.Message {
  getType(): MessageType;
  setType(value: MessageType): void;

  getTime(): number;
  setTime(value: number): void;

  getTopic(): string;
  setTopic(value: string): void;

  getAlias(): string;
  setAlias(value: string): void;

  getBody(): Uint8Array | string;
  getBody_asU8(): Uint8Array;
  getBody_asB64(): string;
  setBody(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TopicMessage.AsObject;
  static toObject(includeInstance: boolean, msg: TopicMessage): TopicMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TopicMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TopicMessage;
  static deserializeBinaryFromReader(message: TopicMessage, reader: jspb.BinaryReader): TopicMessage;
}

export namespace TopicMessage {
  export type AsObject = {
    type: MessageType,
    time: number,
    topic: string,
    alias: string,
    body: Uint8Array | string,
  }
}

export class PositionData extends jspb.Message {
  getPositionX(): number;
  setPositionX(value: number): void;

  getPositionY(): number;
  setPositionY(value: number): void;

  getPositionZ(): number;
  setPositionZ(value: number): void;

  getRotationX(): number;
  setRotationX(value: number): void;

  getRotationY(): number;
  setRotationY(value: number): void;

  getRotationZ(): number;
  setRotationZ(value: number): void;

  getRotationW(): number;
  setRotationW(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PositionData.AsObject;
  static toObject(includeInstance: boolean, msg: PositionData): PositionData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PositionData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PositionData;
  static deserializeBinaryFromReader(message: PositionData, reader: jspb.BinaryReader): PositionData;
}

export namespace PositionData {
  export type AsObject = {
    positionX: number,
    positionY: number,
    positionZ: number,
    rotationX: number,
    rotationY: number,
    rotationZ: number,
    rotationW: number,
  }
}

export class ProfileData extends jspb.Message {
  getAvatarType(): string;
  setAvatarType(value: string): void;

  getDisplayName(): string;
  setDisplayName(value: string): void;

  getPublicKey(): string;
  setPublicKey(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProfileData.AsObject;
  static toObject(includeInstance: boolean, msg: ProfileData): ProfileData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ProfileData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProfileData;
  static deserializeBinaryFromReader(message: ProfileData, reader: jspb.BinaryReader): ProfileData;
}

export namespace ProfileData {
  export type AsObject = {
    avatarType: string,
    displayName: string,
    publicKey: string,
  }
}

export class ChatData extends jspb.Message {
  getMessageId(): string;
  setMessageId(value: string): void;

  getText(): string;
  setText(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatData.AsObject;
  static toObject(includeInstance: boolean, msg: ChatData): ChatData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ChatData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatData;
  static deserializeBinaryFromReader(message: ChatData, reader: jspb.BinaryReader): ChatData;
}

export namespace ChatData {
  export type AsObject = {
    messageId: string,
    text: string,
  }
}

export enum MessageType {
  UNKNOWN_MESSAGE_TYPE = 0,
  WELCOME_SERVER = 1,
  WELCOME_CLIENT = 2,
  CONNECT = 3,
  WEBRTC_SUPPORTED = 4,
  WEBRTC_OFFER = 5,
  WEBRTC_ANSWER = 6,
  WEBRTC_ICE_CANDIDATE = 7,
  PING = 8,
  ADD_TOPIC = 9,
  REMOVE_TOPIC = 10,
  TOPIC = 11,
}

