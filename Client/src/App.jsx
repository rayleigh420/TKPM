import './App.css'
import { Route, Routes } from 'react-router-dom'
import Layout from './pages/Layout'
import Home from './pages/home'
import Book from './pages/Book'

function App() {

  return (
    <>
      <Routes>
        <Route path="/" element={<Layout />} >
          <Route path='home' element={<Home />} />
          <Route path='book' element={<Book />} />
        </Route>
      </Routes>
    </>
  )
}

export default App
