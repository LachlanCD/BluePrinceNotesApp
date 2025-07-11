import './App.css'
import Navbar from './components/Navbar'
import CreationForm from './pages/AddNewPage'
import GeneralHomePage from './pages/GeneralHome'
import RoomHomePage from './pages/RoomHome'
import { Routes, Route } from 'react-router-dom'
import RoomNotePage from './pages/RoomNote'
import GeneralNotePage from './pages/GeneralNote'

function App() {

  return (
    <>
      <Navbar />
      <Routes>
        <Route path='/' element={<RoomHomePage />} />
        <Route path='/generals' element={<GeneralHomePage />} />
        <Route path='/add-new' element={<CreationForm />} />
        <Route path='/rooms/:id' element={<RoomNotePage />} />
        <Route path='/generals/:id' element={<GeneralNotePage />} />
      </Routes>
    </>
  )
}

export default App
