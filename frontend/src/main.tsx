import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
// import './index.css'
// import App from './App.tsx'
import AddTakeAHistory from './pages/TakeAHistory/AddTakeAHistory/index.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AddTakeAHistory />
  </StrictMode>,
)
