import './App.css'
import Navbar from './components/Navbar'
import RoomHomePage from './components/RoomHome'
import { Routes, Route } from 'react-router-dom'

function App() {

  return (
    <>
      <body className="pt-20">
        <Navbar />
        <Routes>
          <Route path='/' element={<RoomHomePage />} />
        </Routes>
      </body>
    </>
  )
}

export default App
