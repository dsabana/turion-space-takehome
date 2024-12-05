import {
    ColumnDef,
    flexRender,
    getCoreRowModel,
    getPaginationRowModel,
    getSortedRowModel,
    SortingState,
    useReactTable,
} from "@tanstack/react-table";
import { ReactElement, useMemo, useState } from "react";
import { HiBarsArrowDown, HiBarsArrowUp } from "react-icons/hi2";
import {
    MdKeyboardDoubleArrowLeft,
    MdKeyboardDoubleArrowRight,
    MdOutlineKeyboardArrowLeft,
    MdOutlineKeyboardArrowRight,
    MdOutlineSubject,
} from "react-icons/md";

type CustomTableProps = {
    columns: ColumnDef<any, any | undefined>[];
    data: any;
    isSortable?: boolean;
    withPagingControls?: boolean;
    tableClassName?: string;
    cellClassName?: string;
};

export function CustomTable({
    columns,
    data,
    isSortable = false,
    withPagingControls = false,
    tableClassName,
    cellClassName,
}: CustomTableProps): ReactElement {
    const finalData = useMemo(() => data, [data]);
    const finalColumnDef = useMemo(() => columns, [columns]);

    const [sorting, setSorting] = useState<SortingState>([]);

    const tableInstance = useReactTable({
        columns: finalColumnDef,
        data: finalData,
        getCoreRowModel: getCoreRowModel(),
        getSortedRowModel: isSortable ? getSortedRowModel() : undefined,
        state: {
            sorting: isSortable ? sorting : undefined,
        },
        onSortingChange: setSorting,
        getPaginationRowModel: withPagingControls
            ? getPaginationRowModel()
            : undefined,
    });

    return (
        <div>
            <section className="p-8 bg-athens-gray-300 text-edward-950 rounded-2xl">
                <table className={`w-full ${tableClassName}`}>
                    <thead className="border-b-2 border-edward-700 font-medium">
                    {tableInstance
                        .getHeaderGroups()
                        .map((headerElement) => {
                            return (
                                <tr key={headerElement.id}>
                                    {headerElement.headers.map(
                                        (columnElement) => {
                                            return (
                                                <th
                                                    key={columnElement.id}
                                                    colSpan={
                                                        columnElement.colSpan
                                                    }
                                                    onClick={columnElement.column.getToggleSortingHandler()}
                                                    className="p-2 pt-0 text-left"
                                                >
                                                    <span className="flex">
                                                        {flexRender(
                                                            columnElement
                                                                .column
                                                                .columnDef
                                                                .header,
                                                            columnElement.getContext(),
                                                        )}
                                                        {isSortable &&
                                                            ({
                                                                asc: (
                                                                    <HiBarsArrowUp className="ml-1 self-center" />
                                                                ),
                                                                desc: (
                                                                    <HiBarsArrowDown className="ml-1 self-center" />
                                                                ),
                                                            }[
                                                                columnElement.column.getIsSorted() as string
                                                                ] ?? (
                                                                <MdOutlineSubject className="ml-1 self-center" />
                                                            ))}
                                                    </span>
                                                </th>
                                            );
                                        },
                                    )}
                                </tr>
                            );
                        })}
                    </thead>
                    <tbody>
                    {tableInstance.getRowModel().rows.map((rowElement) => {
                        return (
                            <tr
                                key={rowElement.id}
                                className="border-b-2 border-edward-700 text-edward-950"
                            >
                                {rowElement
                                    .getVisibleCells()
                                    .map((cellElement) => {
                                        return (
                                            <td
                                                key={cellElement.id}
                                                className={`px-2 py-5 ${cellClassName}`}
                                            >
                                                {flexRender(
                                                    cellElement.column
                                                        .columnDef.cell,
                                                    cellElement.getContext(),
                                                )}
                                            </td>
                                        );
                                    })}
                            </tr>
                        );
                    })}
                    </tbody>
                </table>
            </section>
            {withPagingControls && (
                <section className="py-3 px-6 mt-2 flex items-center justify-between bg-athens-gray-950 rounded-2xl text-white">
                    <section className="flex items-center gap-5 text-black">
                        <select
                            className="p-2 border-2 bg-white rounded-md"
                            value={tableInstance.getState().pagination.pageSize}
                            onChange={(e) => {
                                tableInstance.setPageSize(
                                    Number(e.target.value),
                                );
                            }}
                        >
                            {[10, 20, 30, 40, 50].map((pageSize) => (
                                <option key={pageSize} value={pageSize}>
                                    {pageSize}
                                </option>
                            ))}
                        </select>
                        <span className={"text-white"}>Items per page</span>
                    </section>

                    <section className="flex gap-2 items-center">
                        <span className="flex items-center gap-1">
                            <div>Page</div>
                            <strong>
                                {tableInstance.getState().pagination.pageIndex +
                                    1}{" "}
                                of {tableInstance.getPageCount()}
                            </strong>
                        </span>
                        <button
                            className={`p-1 ${
                                tableInstance.getCanPreviousPage()
                                    ? "text-black"
                                    : "text-slate-400"
                            }`}
                            onClick={() => tableInstance.setPageIndex(0)}
                            disabled={!tableInstance.getCanPreviousPage()}
                        >
                            {
                                <MdKeyboardDoubleArrowLeft
                                    color={"white"}
                                    size={25}
                                />
                            }
                        </button>
                        <button
                            className={`p-1 ${
                                tableInstance.getCanPreviousPage()
                                    ? "text-black"
                                    : "text-slate-400"
                            }`}
                            onClick={() => tableInstance.previousPage()}
                            disabled={!tableInstance.getCanPreviousPage()}
                        >
                            {
                                <MdOutlineKeyboardArrowLeft
                                    color={"white"}
                                    size={25}
                                />
                            }
                        </button>
                        <button
                            className={`p-1 ${
                                tableInstance.getCanNextPage()
                                    ? "text-black"
                                    : "text-slate-400"
                            }`}
                            onClick={() => tableInstance.nextPage()}
                            disabled={!tableInstance.getCanNextPage()}
                        >
                            {
                                <MdOutlineKeyboardArrowRight
                                    color={"white"}
                                    size={25}
                                />
                            }
                        </button>
                        <button
                            className={`p-1 ${
                                tableInstance.getCanNextPage()
                                    ? "text-black"
                                    : "text-slate-400"
                            }`}
                            onClick={() =>
                                tableInstance.setPageIndex(
                                    tableInstance.getPageCount() - 1,
                                )
                            }
                            disabled={!tableInstance.getCanNextPage()}
                        >
                            {
                                <MdKeyboardDoubleArrowRight
                                    color={"white"}
                                    size={25}
                                />
                            }
                        </button>
                    </section>
                </section>
            )}
        </div>
    );
}
