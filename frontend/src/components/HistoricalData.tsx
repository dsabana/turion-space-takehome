import {ReactElement} from "react";
import {TelemetryPacket} from "../types.ts";
import axios from "axios";
import {CustomTable} from "./CustomTable.tsx";
import {createColumnHelper} from "@tanstack/react-table";
import {formatNumber} from "../utils.ts";
import {useQuery} from "react-query";

const fetchHistoricalData = async (): Promise<TelemetryPacket[]> => {
    const { data } = await axios.get<TelemetryPacket[]>(`http://localhost:8080/api/v1/telemetry`);
    return data;
}

export function HistoricalData(): ReactElement {
    const { data: packetList, isSuccess, refetch } = useQuery(['historicalData'], fetchHistoricalData);

    const columnHelper = createColumnHelper<TelemetryPacket>();
    const columns = [
        columnHelper.accessor("primaryHeader.packetId", {
            header: () => "Packet ID",
            cell: (info) => info.getValue(),
        }),
        columnHelper.accessor("primaryHeader.packetSeqCtrl", {
            header: () => "Packet Seq Ctrl",
            cell: (info) => info.getValue(),
        }),
        columnHelper.accessor("secondaryHeader.timestamp", {
            header: () => "Timestamp",
            cell: (info) => info.getValue(),
        }),
        columnHelper.accessor("payload.battery", {
            header: () => "Battery",
            cell: (info) => `${formatNumber(info.getValue(), 2)}%`,
        }),
        columnHelper.accessor("payload.altitude", {
            header: () => "Altitude",
            cell: (info) => `${formatNumber(info.getValue(), 2)}km`,
        }),
        columnHelper.accessor("payload.signal", {
            header: () => "Signal",
            cell: (info) => `${formatNumber(info.getValue(), 2)}dB`,
        }),
        columnHelper.accessor("payload.temperature", {
            header: () => "Temperature",
            cell: (info) => `${formatNumber(info.getValue(), 2)}°C`,
        }),
    ]

    return (
        <div className={"bg-gray-400 rounded-2xl w-[1024px] p-10"}>
            <div className={"flex text-left mb-8"}>
                <h2 className={"font-bold grow text-3xl"}>Historical Telemetry Data</h2>
                <button
                    className={"bg-edward-950 hover:bg-edward-900 active:bg-edward-700 text-white p-4 rounded-xl font-semibold"}
                    onClick={() => refetch()}
                >
                    Refetch Data
                </button>
            </div>

            {isSuccess && packetList?.length && (
                <CustomTable
                    columns={columns}
                    data={packetList}
                    withPagingControls
                    isSortable
                />
            )}
        </div>
    )
}