import {ReactElement} from "react";
import {TelemetryPacket} from "../types.ts";
import axios from "axios";
import {formatNumber} from "../utils.ts";
import { useQuery } from "react-query";

const fetchCurrentData = async (): Promise<TelemetryPacket> => {
    const { data } = await axios.get<TelemetryPacket>(`http://localhost:8080/api/v1/telemetry/current`);
    return data;
}

export function CurrentData(): ReactElement {
    const { data: packet } = useQuery(['currentData'], fetchCurrentData, {
        keepPreviousData: true,
        refetchInterval: 1000,
    });

    return (
        <div className={`w-[480px] ${packet?.hasAnomaly ? "bg-orange-500" : "bg-gray-200"} rounded-2xl p-10`}>
            <div className={"mb-5"}>
                <h1 className={"text-3xl font-bold"}>Current Telemetry Data</h1>
                <h2 className={"text-xl font-bold"}>Updated in Real Time</h2>
            </div>

            <div className={"grid grid-cols-2 gap-2"}>
                <label className={"font-semibold"}>Packet ID</label>
                <label>{packet?.primaryHeader.packetId}</label>

                <label className={"font-semibold"}>Packet Seq Ctrl</label>
                <label>{packet?.primaryHeader.packetSeqCtrl}</label>

                <label className={"font-semibold"}>Packet Length</label>
                <label>{packet?.primaryHeader.packetLength}</label>

                <label className={"font-semibold"}>Subsystem ID</label>
                <label>{packet?.secondaryHeader.subsystemId}</label>

                <label className={"font-semibold"}>Timestamp</label>
                <label>{packet?.secondaryHeader.timestamp}</label>

                <label className={"font-semibold"}>Altitude</label>
                <label>{formatNumber(packet?.payload.altitude ?? 0, 2)}</label>

                <label className={"font-semibold"}>Battery</label>
                <label>{formatNumber(packet?.payload.battery ?? 0, 2)}</label>

                <label className={"font-semibold"}>Signal</label>
                <label>{formatNumber(packet?.payload.signal ?? 0, 2)}</label>

                <label className={"font-semibold"}>Temperature</label>
                <label>{formatNumber(packet?.payload.temperature ?? 0, 2)}</label>

                <label className={"col-span-2 mt-8 mx-auto text-2xl font-black"}>{packet?.hasAnomaly ? "ANOMALY DETECTED" : "NORMAL DATA"}</label>
            </div>
        </div>
    )
}