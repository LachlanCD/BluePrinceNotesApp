import './App.css'
import CreationForm from './pages/AddNewPage'
import GeneralHomePage from './pages/GeneralHome'
import RoomHomePage from './pages/RoomHome'
import { Routes, Route } from 'react-router-dom'
import RoomNotePage from './pages/RoomNote'
import GeneralNotePage from './pages/GeneralNote'
import PageLayout from './pages/PageLayout'

function App() {

  return (
    <>
      <Routes>
        <Route path='/' element={<RoomHomePage />} />
        <Route path='/:workspaceID' element={<PageLayout />} >
          <Route path='/:workspaceID/rooms' element={<RoomHomePage />} />
          <Route path='/:workspaceID/generals' element={<GeneralHomePage />} />
          <Route path='/:workspaceID/add-new' element={<CreationForm />} />
          <Route path='/:workspaceID/rooms/:id' element={<RoomNotePage />} />
          <Route path='/:workspaceID/generals/:id' element={<GeneralNotePage />} />
        </Route>
      </Routes>
    </>
  )
}

export default App
