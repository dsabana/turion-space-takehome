import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import {QueryClient, QueryClientProvider} from "react-query";
import {App} from "./App.tsx";

const queryClient = new QueryClient();

createRoot(document.getElementById('root')!).render(
  <StrictMode>
      <QueryClientProvider client={queryClient} >
        <App />
      </QueryClientProvider>
  </StrictMode>,
)