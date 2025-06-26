import './App.css'
import Navbar from './components/Navbar'
import CreationForm from './pages/AddNewPage'
import GeneralHomePage from './pages/GeneralHome'
import RoomHomePage from './pages/RoomHome'
import { Routes, Route } from 'react-router-dom'

function App() {

  return (
    <div className="pt-20">
      <Navbar />
      <Routes>
        <Route path='/' element={<RoomHomePage />} />
        <Route path='/general' element={<GeneralHomePage />} />
        <Route path='/add-new' element={<CreationForm />} />
      </Routes>
    </div>
  )
}

export default App
