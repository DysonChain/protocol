import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "names";
export interface MsgRegister {
    owner: string;
    commit: string;
    price: string;
}
export interface MsgRegisterResponse {
}
export interface MsgCreateName {
    owner: string;
    name: string;
    destination: string;
    price: string;
    expires: string;
    authorized: string;
    commit: string;
    salt: string;
}
export interface MsgCreateNameResponse {
}
export interface MsgUpdateName {
    owner: string;
    name: string;
    destination: string;
    price: string;
    expires: string;
    authorized: string;
    commit: string;
    salt: string;
}
export interface MsgUpdateNameResponse {
}
export interface MsgDeleteName {
    owner: string;
    name: string;
}
export interface MsgDeleteNameResponse {
}
export interface MsgReveal {
    creator: string;
    name: string;
    salt: string;
}
export interface MsgRevealResponse {
}
export declare const MsgRegister: {
    encode(message: MsgRegister, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegister;
    fromJSON(object: any): MsgRegister;
    toJSON(message: MsgRegister): unknown;
    fromPartial(object: DeepPartial<MsgRegister>): MsgRegister;
};
export declare const MsgRegisterResponse: {
    encode(_: MsgRegisterResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterResponse;
    fromJSON(_: any): MsgRegisterResponse;
    toJSON(_: MsgRegisterResponse): unknown;
    fromPartial(_: DeepPartial<MsgRegisterResponse>): MsgRegisterResponse;
};
export declare const MsgCreateName: {
    encode(message: MsgCreateName, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateName;
    fromJSON(object: any): MsgCreateName;
    toJSON(message: MsgCreateName): unknown;
    fromPartial(object: DeepPartial<MsgCreateName>): MsgCreateName;
};
export declare const MsgCreateNameResponse: {
    encode(_: MsgCreateNameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateNameResponse;
    fromJSON(_: any): MsgCreateNameResponse;
    toJSON(_: MsgCreateNameResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateNameResponse>): MsgCreateNameResponse;
};
export declare const MsgUpdateName: {
    encode(message: MsgUpdateName, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateName;
    fromJSON(object: any): MsgUpdateName;
    toJSON(message: MsgUpdateName): unknown;
    fromPartial(object: DeepPartial<MsgUpdateName>): MsgUpdateName;
};
export declare const MsgUpdateNameResponse: {
    encode(_: MsgUpdateNameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateNameResponse;
    fromJSON(_: any): MsgUpdateNameResponse;
    toJSON(_: MsgUpdateNameResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateNameResponse>): MsgUpdateNameResponse;
};
export declare const MsgDeleteName: {
    encode(message: MsgDeleteName, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteName;
    fromJSON(object: any): MsgDeleteName;
    toJSON(message: MsgDeleteName): unknown;
    fromPartial(object: DeepPartial<MsgDeleteName>): MsgDeleteName;
};
export declare const MsgDeleteNameResponse: {
    encode(_: MsgDeleteNameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteNameResponse;
    fromJSON(_: any): MsgDeleteNameResponse;
    toJSON(_: MsgDeleteNameResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteNameResponse>): MsgDeleteNameResponse;
};
export declare const MsgReveal: {
    encode(message: MsgReveal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgReveal;
    fromJSON(object: any): MsgReveal;
    toJSON(message: MsgReveal): unknown;
    fromPartial(object: DeepPartial<MsgReveal>): MsgReveal;
};
export declare const MsgRevealResponse: {
    encode(_: MsgRevealResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRevealResponse;
    fromJSON(_: any): MsgRevealResponse;
    toJSON(_: MsgRevealResponse): unknown;
    fromPartial(_: DeepPartial<MsgRevealResponse>): MsgRevealResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    Register(request: MsgRegister): Promise<MsgRegisterResponse>;
    CreateName(request: MsgCreateName): Promise<MsgCreateNameResponse>;
    UpdateName(request: MsgUpdateName): Promise<MsgUpdateNameResponse>;
    DeleteName(request: MsgDeleteName): Promise<MsgDeleteNameResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    Reveal(request: MsgReveal): Promise<MsgRevealResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    Register(request: MsgRegister): Promise<MsgRegisterResponse>;
    CreateName(request: MsgCreateName): Promise<MsgCreateNameResponse>;
    UpdateName(request: MsgUpdateName): Promise<MsgUpdateNameResponse>;
    DeleteName(request: MsgDeleteName): Promise<MsgDeleteNameResponse>;
    Reveal(request: MsgReveal): Promise<MsgRevealResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
