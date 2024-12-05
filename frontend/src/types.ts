export type TelemetryPacket = {
    primaryHeader: PrimaryHeader;
    secondaryHeader: SecondaryHeader;
    payload: TelemetryPayload;
    hasAnomaly: boolean;
}

export type PrimaryHeader = {
    packetId: number;
    packetSeqCtrl: number;
    packetLength: number;
}

export type SecondaryHeader = {
    subsystemId: number;
    timestamp: string;
}

export type TelemetryPayload = {
    altitude: number;
    battery: number;
    signal: number;
    temperature: number;
}