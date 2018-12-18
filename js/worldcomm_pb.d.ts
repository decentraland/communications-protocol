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

export class PositionMessage extends jspb.Message {
  getType(): WorldCommMessageType;
  setType(value: WorldCommMessageType): void;

  getTime(): number;
  setTime(value: number): void;

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

  getAlias(): number;
  setAlias(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PositionMessage.AsObject;
  static toObject(includeInstance: boolean, msg: PositionMessage): PositionMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PositionMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PositionMessage;
  static deserializeBinaryFromReader(message: PositionMessage, reader: jspb.BinaryReader): PositionMessage;
}

export namespace PositionMessage {
  export type AsObject = {
    type: WorldCommMessageType,
    time: number,
    positionX: number,
    positionY: number,
    positionZ: number,
    rotationX: number,
    rotationY: number,
    rotationZ: number,
    rotationW: number,
    alias: number,
  }
}

export class ProfileMessage extends jspb.Message {
  getType(): WorldCommMessageType;
  setType(value: WorldCommMessageType): void;

  getTime(): number;
  setTime(value: number): void;

  getPositionX(): number;
  setPositionX(value: number): void;

  getPositionZ(): number;
  setPositionZ(value: number): void;

  getAvatarType(): string;
  setAvatarType(value: string): void;

  getDisplayName(): string;
  setDisplayName(value: string): void;

  getPeerId(): string;
  setPeerId(value: string): void;

  getAlias(): number;
  setAlias(value: number): void;

  getPublicKey(): string;
  setPublicKey(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProfileMessage.AsObject;
  static toObject(includeInstance: boolean, msg: ProfileMessage): ProfileMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ProfileMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProfileMessage;
  static deserializeBinaryFromReader(message: ProfileMessage, reader: jspb.BinaryReader): ProfileMessage;
}

export namespace ProfileMessage {
  export type AsObject = {
    type: WorldCommMessageType,
    time: number,
    positionX: number,
    positionZ: number,
    avatarType: string,
    displayName: string,
    peerId: string,
    alias: number,
    publicKey: string,
  }
}

export class ChatMessage extends jspb.Message {
  getType(): WorldCommMessageType;
  setType(value: WorldCommMessageType): void;

  getTime(): number;
  setTime(value: number): void;

  getMessageId(): string;
  setMessageId(value: string): void;

  getPositionX(): number;
  setPositionX(value: number): void;

  getPositionZ(): number;
  setPositionZ(value: number): void;

  getText(): string;
  setText(value: string): void;

  getAlias(): number;
  setAlias(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatMessage.AsObject;
  static toObject(includeInstance: boolean, msg: ChatMessage): ChatMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ChatMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatMessage;
  static deserializeBinaryFromReader(message: ChatMessage, reader: jspb.BinaryReader): ChatMessage;
}

export namespace ChatMessage {
  export type AsObject = {
    type: WorldCommMessageType,
    time: number,
    messageId: string,
    positionX: number,
    positionZ: number,
    text: string,
    alias: number,
  }
}

export class ClientDisconnectedFromServerMessage extends jspb.Message {
  getType(): WorldCommMessageType;
  setType(value: WorldCommMessageType): void;

  getTime(): number;
  setTime(value: number): void;

  getAlias(): number;
  setAlias(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientDisconnectedFromServerMessage.AsObject;
  static toObject(includeInstance: boolean, msg: ClientDisconnectedFromServerMessage): ClientDisconnectedFromServerMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClientDisconnectedFromServerMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientDisconnectedFromServerMessage;
  static deserializeBinaryFromReader(message: ClientDisconnectedFromServerMessage, reader: jspb.BinaryReader): ClientDisconnectedFromServerMessage;
}

export namespace ClientDisconnectedFromServerMessage {
  export type AsObject = {
    type: WorldCommMessageType,
    time: number,
    alias: number,
  }
}

export class ClockSkewMessage extends jspb.Message {
  getType(): WorldCommMessageType;
  setType(value: WorldCommMessageType): void;

  getTime(): number;
  setTime(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClockSkewMessage.AsObject;
  static toObject(includeInstance: boolean, msg: ClockSkewMessage): ClockSkewMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClockSkewMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClockSkewMessage;
  static deserializeBinaryFromReader(message: ClockSkewMessage, reader: jspb.BinaryReader): ClockSkewMessage;
}

export namespace ClockSkewMessage {
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

export enum WorldCommMessageType {
  UNKNOWN_MESSAGE_TYPE = 0,
  POSITION = 2,
  CHAT = 3,
  CLIENT_DISCONNECTED_FROM_SERVER = 4,
  PROFILE = 5,
  CLOCK_SKEW_DETECTED = 6,
  PING = 8,
}

