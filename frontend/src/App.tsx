import './App.css'
import {CurrentData} from "./components/CurrentData.tsx";
import {ReactElement} from "react";
import {HistoricalData} from "./components/HistoricalData.tsx";

export function App(): ReactElement {
    return (
        <div className={"grid grid-cols-1 gap-10 justify-items-center my-20"}>
            <CurrentData />
            <HistoricalData/>
        </div>
    )
}
