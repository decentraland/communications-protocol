// package: 
// file: worldcomm.proto

import * as jspb from "google-protobuf";

export class WorldCommMessage extends jspb.Message {
  getType(): WorldCommMessageType;
  setType(value: WorldCommMessageType): void;

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
    type: WorldCommMessageType,
    time: number,
  }
}

export class PingMessage extends jspb.Message {
  getType(): WorldCommMessageType;
  setType(value: WorldCommMessageType): void;

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
    type: WorldCommMessageType,
    time: number,
  }
}

export class AddTopicMessage extends jspb.Message {
  getType(): WorldCommMessageType;
  setType(value: WorldCommMessageType): void;

  getTime(): number;
  setTime(value: number): void;

  getTopic(): string;
  setTopic(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddTopicMessage.AsObject;
  static toObject(includeInstance: boolean, msg: AddTopicMessage): AddTopicMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddTopicMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddTopicMessage;
  static deserializeBinaryFromReader(message: AddTopicMessage, reader: jspb.BinaryReader): AddTopicMessage;
}

export namespace AddTopicMessage {
  export type AsObject = {
    type: WorldCommMessageType,
    time: number,
    topic: string,
  }
}

export class RemoveTopicMessage extends jspb.Message {
  getType(): WorldCommMessageType;
  setType(value: WorldCommMessageType): void;

  getTime(): number;
  setTime(value: number): void;

  getTopic(): string;
  setTopic(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveTopicMessage.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveTopicMessage): RemoveTopicMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RemoveTopicMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveTopicMessage;
  static deserializeBinaryFromReader(message: RemoveTopicMessage, reader: jspb.BinaryReader): RemoveTopicMessage;
}

export namespace RemoveTopicMessage {
  export type AsObject = {
    type: WorldCommMessageType,
    time: number,
    topic: string,
  }
}

export class TopicMessage extends jspb.Message {
  getType(): WorldCommMessageType;
  setType(value: WorldCommMessageType): void;

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
    type: WorldCommMessageType,
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

export enum WorldCommMessageType {
  UNKNOWN_MESSAGE_TYPE = 0,
  PING = 1,
  ADD_TOPIC = 2,
  REMOVE_TOPIC = 3,
  TOPIC = 4,
}

