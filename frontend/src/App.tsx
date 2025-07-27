import './App.css'
import CreationForm from './pages/AddNewPage'
import GeneralHomePage from './pages/GeneralHome'
import RoomHomePage from './pages/RoomHome'
import { Routes, Route } from 'react-router-dom'
import RoomNotePage from './pages/RoomNote'
import GeneralNotePage from './pages/GeneralNote'
import PageLayout from './pages/PageLayout'
import NotFound from './pages/NotFoundPage'
import HomePage from './pages/Home'

function App() {

  return (
    <>
      <Routes>
        <Route path='/' element={<HomePage />} />
        <Route path='/:workspaceID' element={<PageLayout />} >
          <Route index element={<HomePage />} />
          <Route path='rooms' element={<RoomHomePage />} />
          <Route path='generals' element={<GeneralHomePage />} />
          <Route path='add-new' element={<CreationForm />} />
          <Route path='rooms/:id' element={<RoomNotePage />} />
          <Route path='generals/:id' element={<GeneralNotePage />} />
          <Route path='*' element={<NotFound />} />
        </Route>
      </Routes>
    </>
  )
}

export default App
