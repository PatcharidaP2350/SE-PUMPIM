import { App } from 'antd'
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import AddTakeAHistory from './pages/TakeAHistory/AddTakeAHistory/takeahistory1'
import AddTakeAHistory2 from './pages/TakeAHistory/AddTakeAHistory2/takeahistory2'
import SaveTakeAHistory from './pages/TakeAHistory/SaveTakeAHistory/saveahistory'
import UpdateTakeAHistory from './pages/TakeAHistory/UpdateTakeAHistory/UpdateTakeAHistory'
import AddPatient from './pages/TakeAHistory/CreatePatient/AddPatient'
// import './index.css'
// import App from './App.tsx'



createRoot(document.getElementById('root')!).render(
  <StrictMode>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             
    <AddPatient/>
  </StrictMode>
)
